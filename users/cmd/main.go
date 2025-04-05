package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	`sync`
	"syscall"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/golang-migrate/migrate/v4"
	postgresMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	`github.com/grpc-ecosystem/grpc-gateway/v2/runtime`
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"lms-user/cmd/config"
	"lms-user/internal/adapters/postgres"
	"lms-user/internal/adapters/storage"
	"lms-user/internal/adapters/storage/cloud"
	"lms-user/internal/adapters/storage/local"
	"lms-user/internal/app"
	"lms-user/internal/metrics"
	grpcPort "lms-user/internal/ports/grpc"
	"lms-user/internal/ports/httpgin"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	metric, err := metrics.NewUserOpenTelemetryMetric()
	if err != nil {
		log.Fatalf("don`t create metric %v", err)
	}

	// initial domain layer.
	domainUser := app.CreateAppUser(
		newPostgres(), newStorage(), metric,
	)

	// Create context that listens for interrupt signals
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Create errgroup with the signal context
	eg, ctx := errgroup.WithContext(ctx)

	// Initialize servers
	httpServer := &http.Server{
		Addr:    ":18080",
		Handler: httpgin.Server(":18080", domainUser).Handler,
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcPort.NewService(domainUser).Interceptor(),
		),
	)
	grpcPort.RegisterUserServiceServer(grpcServer, grpcPort.NewService(domainUser))

	// gRPC Gateway setup
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = grpcPort.RegisterUserServiceHandlerFromEndpoint(
		context.Background(), // Use separate context for setup
		mux,
		"localhost:50054",
		opts,
	)
	if err != nil {
		log.Fatal(err)
	}

	gatewayServer := &http.Server{
		Addr:    ":5051",
		Handler: mux,
	}

	// Start servers in separate goroutines
	var wg sync.WaitGroup

	// gRPC Server
	wg.Add(1)
	eg.Go(func() error {
		defer wg.Done()

		lis, err := net.Listen("tcp", ":50054")
		if err != nil {
			return fmt.Errorf("failed to listen: %v", err)
		}

		log.Printf("starting gRPC server on %s", lis.Addr().String())
		defer log.Printf("stopping gRPC server")

		errCh := make(chan error, 1)
		go func() {
			if err := grpcServer.Serve(lis); err != nil {
				errCh <- fmt.Errorf("gRPC server error: %w", err)
			}
		}()

		select {
		case <-ctx.Done():
			grpcServer.GracefulStop()
			return nil
		case err := <-errCh:
			return err
		}
	})

	// HTTP Server
	wg.Add(1)
	eg.Go(func() error {
		defer wg.Done()

		log.Printf("starting HTTP server on %s", httpServer.Addr)
		defer log.Printf("stopping HTTP server")

		errCh := make(chan error, 1)
		go func() {
			if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				errCh <- fmt.Errorf("HTTP server error: %w", err)
			}
		}()

		select {
		case <-ctx.Done():
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			if err := httpServer.Shutdown(shutdownCtx); err != nil {
				return fmt.Errorf("HTTP server shutdown error: %w", err)
			}
			return nil
		case err := <-errCh:
			return err
		}
	})

	// gRPC Gateway Server
	wg.Add(1)
	eg.Go(func() error {
		defer wg.Done()

		log.Printf("starting gRPC-Gateway server on %s", gatewayServer.Addr)
		defer log.Printf("stopping gRPC-Gateway server")

		errCh := make(chan error, 1)
		go func() {
			if err := gatewayServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				errCh <- fmt.Errorf("gRPC-Gateway server error: %w", err)
			}
		}()

		select {
		case <-ctx.Done():
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			if err := gatewayServer.Shutdown(shutdownCtx); err != nil {
				return fmt.Errorf("gRPC-Gateway server shutdown error: %w", err)
			}
			return nil
		case err := <-errCh:
			return err
		}
	})

	// Wait for shutdown signal
	eg.Go(func() error {
		<-ctx.Done()
		log.Println("received shutdown signal")
		return nil
	})

	// Wait for all servers to shutdown
	if err := eg.Wait(); err != nil {
		log.Printf("error in server group: %v", err)
	}

	// Additional wait to ensure all cleanup is done
	wg.Wait()
	log.Println("all servers stopped gracefully")
}

func newStorage() storage.IFace {
	ymlAWS, err := config.NewCfgAWS()
	if errors.Is(err, os.ErrNotExist) {
		return local.NewLocal()
	} else if err != nil {
		log.Fatalf("config read error %e", err)
	}

	awsS3, err := cloud.NewAWS(ymlAWS.AWS)
	if err != nil {
		log.Fatalf("s3session error %e obj: %v", err, awsS3)
	}

	return awsS3
}

func newPostgres() postgres.IUserPostgres {
	ymlPostgres, err := config.NewCfgPostgres()
	if err != nil {
		log.Fatalf("postgres yml error %e obj: %v", err, ymlPostgres)
	}
	cfg := postgres.Config{
		Host:     ymlPostgres.Postgres.Host,
		Port:     ymlPostgres.Postgres.Port,
		User:     ymlPostgres.Postgres.User,
		DBName:   ymlPostgres.Postgres.DBName,
		Password: ymlPostgres.Postgres.Password,
		SSLmode:  ymlPostgres.Postgres.SSLmode,
	}

	db, err := postgres.CreatePostgres(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)

	if err = runMigrations(db); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	return postgres.CreateRepoUser(db)
}

func runMigrations(db *sqlx.DB) error {
	driver, err := postgresMigrate.WithInstance(db.DB, &postgresMigrate.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}
