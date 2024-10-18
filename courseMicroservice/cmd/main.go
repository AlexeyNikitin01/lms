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
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	nosql "course/internal/adapters/mongo"
	grpcPort "course/internal/ports/grpc"

	"course/cmd/config"
	"course/internal/adapters/postgres"
	"course/internal/app"
	"course/internal/ports/httpgin"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	conn := ConnectToDataBase()

	domainCourse := app.NewCourseApp(
		postgres.CreateRepoUser(conn.DBConn),
		nosql.NewMongoRepo(conn.MongoConn),
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
