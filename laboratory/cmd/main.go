package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	`sync`
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	postgresMigrate "github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/friendsofgo/errors"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"lab/cmd/config"
	"lab/internal/adapters/postgres"
	"lab/internal/app"
	"lab/internal/ports/http"
	user "lab/pkg/grpc"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg, err := config.NewLabCourseMicroservice()
	if err != nil {
		logrus.Fatal(err)
	}
	db, err := ConnPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err = runMigrations(db); err != nil {
		logrus.Fatal(err)
	}
	domainUser := app.NewLab(postgres.NewLabPostgres(db))

	// Create context that listens for interrupt signals
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Create errgroup with the signal context
	eg, ctx := errgroup.WithContext(ctx)

	userConn, err := grpc.NewClient(":50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to user service: %v", err)
	}
	defer userConn.Close()

	userClient := user.NewUserServiceClient(userConn)

	// Initialize servers
	httpServer := &http.Server{
		Addr:    ":3344",
		Handler: httpgin.Server(":3344", domainUser, userClient).Handler,
	}

	// Start servers in separate goroutines
	var wg sync.WaitGroup

	// HTTP Server
	wg.Add(1)
	eg.Go(func() error {
		defer wg.Done()

		log.Printf("starting HTTP server on %s", httpServer.Addr)
		defer log.Printf("stopping HTTP server")

		errCh := make(chan error, 1)
		go func() {
			if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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

func ConnPostgres(um *config.LabMic) (*sqlx.DB, error) {
	cfg := postgres.Config{
		Host:     um.Postgres.Host,
		Port:     um.Postgres.Port,
		User:     um.Postgres.User,
		DBName:   um.Postgres.DBName,
		Password: um.Postgres.Password,
		SSLmode:  um.Postgres.SSLmode,
	}

	db, err := postgres.CreatePostgres(&cfg)
	if err != nil {
		return nil, err
	}

	boil.SetDB(db)
	boil.DebugMode = true

	return db, nil
}

func runMigrations(db *sqlx.DB) error {
	driver, err := postgresMigrate.WithInstance(db.DB, &postgresMigrate.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
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
