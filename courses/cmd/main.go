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
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"course/cmd/config"
	nosql "course/internal/adapters/mongo"
	"course/internal/adapters/postgres"
	`course/internal/adapters/storage`
	`course/internal/adapters/storage/cloud`
	`course/internal/adapters/storage/local`
	"course/internal/app"
	grpcPort "course/internal/ports/grpc"
	"course/internal/ports/httpgin"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	conn := ConnectToDataBase()

	if err := runMigrations(conn.DBConn); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	domainCourse := app.NewCourseApp(
		postgres.CreateRepoUser(conn.DBConn),
		nosql.NewMongoRepo(conn.MongoConn),
		newStorage(),
	)

	svr := httpgin.Server(":1818", domainCourse)

	httpServer := &http.Server{
		Addr:    ":1818",
		Handler: svr.Handler,
	}

	port := ":50057"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	svc := grpcPort.NewService(
		domainCourse,
	)

	grpcServer := grpc.NewServer()

	grpcPort.RegisterCourseServiceServer(grpcServer, svc)

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
		log.Printf("starting http server, listening on %s\n", ":1818")
		defer log.Printf("close http server listening on %s\n", ":1818")

		errCh := make(chan error)

		defer func() {
			shCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			if err := httpServer.Shutdown(shCtx); err != nil {
				log.Printf("can't close http server listening on %s: %s", ":1818", err.Error())
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

type Conn struct {
	DBConn    *sqlx.DB
	MongoConn *mongo.Client
}

func ConnectToDataBase() *Conn {
	um, err := config.NewConfigCourseMicroservice()
	if err != nil {
		log.Fatalf("config read error %e", err)
	}

	dbConn, err := ConnPostgres(um)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient, err := ConnMongo(um)
	if err != nil {
		log.Fatal(err)
	}

	return &Conn{
		DBConn:    dbConn,
		MongoConn: mongoClient,
	}
}

func ConnPostgres(um *config.CourseMicroservice) (*sqlx.DB, error) {
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

func ConnMongo(um *config.CourseMicroservice) (*mongo.Client, error) {
	cfgMongo := nosql.MongoConfig{
		Host:     um.Mongo.Host,
		Port:     um.Mongo.Port,
		User:     um.Mongo.User,
		Password: um.Mongo.Password,
	}
	mongoClient, err := nosql.NewMongoClient(&cfgMongo)
	if err != nil {
		return nil, err
	}

	return mongoClient, nil
}

func newStorage() storage.ICloud {
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
