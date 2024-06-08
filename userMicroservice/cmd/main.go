package main

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"edu-material/userMicroservice/internal/adapters/postgres"
	"edu-material/userMicroservice/internal/app"
	grpcPort "edu-material/userMicroservice/internal/ports/grpc"
	"edu-material/userMicroservice/internal/ports/httpgin"
)

func main() {
	fmt.Println("init server")

	//cfg := postgres.Config{
	//	Host:     "localhost",
	//	Port:     "7777",
	//	User:     "postgres",
	//	DBName:   "user",
	//	Password: "pass",
	//	SSLmode:  "disable",
	//}

	//a := app.CreateAppUser(postgres.CreateRepoUser(db))
	//
	//s := httpgin.Server(":18080", a)

	//log.Fatal(s.ListenAndServe())

	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg := postgres.Config{
		Host:     "localhost",
		Port:     "7777",
		User:     "postgres",
		DBName:   "user",
		Password: "pass",
		SSLmode:  "disable",
	}

	db, err := postgres.CreatePostgres(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)

	a := app.CreateAppUser(postgres.CreateRepoUser(db))

	svr := httpgin.Server(":18080", a)

	httpServer := &http.Server{
		Addr:    ":18080",
		Handler: svr.Handler,
	}

	port := ":50054"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	svc := grpcPort.NewService()
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

	// run grpc server
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
	//Run rest
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
