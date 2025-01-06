package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/golang-migrate/migrate/v4"
	postgresMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	svr := httpgin.Server(":18080", domainUser)

	httpServer := &http.Server{
		Addr:    ":18080",
		Handler: svr.Handler,
	}

	port := ":50054"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	svc := grpcPort.NewService(
		domainUser,
	)

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(svc.Interceptor()))

	grpcPort.RegisterUserServiceServer(grpcServer, svc)

	eg, ctx := errgroup.WithContext(context.Background())

	sigQuit := make(chan os.Signal, 1)

	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)

	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	eg.Go(func() error {
		select {
		case s := <-sigQuit:
			log.Printf("captured signal: %v\n", s)
			return fmt.Errorf("captured signal: %v", s)
		case <-ctx.Done():
			return nil
		}
	})

	// run grpc server.
	eg.Go(func() error {
		log.Printf("starting grpc server, listening on %s\n", port)
		defer log.Printf("close grpc server listening on %s\n", port)

		errCh := make(chan error)

		defer func() {
			grpcServer.GracefulStop()
			_ = lis.Close()

			close(errCh)
		}()

		go func() {
			if err := grpcServer.Serve(lis); err != nil {
				errCh <- err
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return fmt.Errorf("grpc server can't listen and serve requests: %w", err)
		}
	})
	// Run rest.
	eg.Go(func() error {
		log.Printf("starting http server, listening on %s\n", ":18080")
		defer log.Printf("close http server listening on %s\n", ":18080")

		errCh := make(chan error)

		defer func() {
			shCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			if err := httpServer.Shutdown(shCtx); err != nil {
				log.Printf("can't close http server listening on %s: %s", ":18080", err.Error())
			}

			close(errCh)
		}()

		go func() {
			if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				errCh <- err
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return fmt.Errorf("http server can't listen and serve requests: %w", err)
		}
	})

	if err := eg.Wait(); err != nil {
		log.Printf("gracefully shutting down the servers: %s\n", err.Error())
	}

	log.Println("servers were successfully shutdown")
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
