package testutil

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc/test/bufconn"

	"github.com/AlexeyNikitin01/lms-user/internal/adapters/postgres"
	"github.com/AlexeyNikitin01/lms-user/internal/app"
	grpcPort "github.com/AlexeyNikitin01/lms-user/internal/ports/grpc"
	"github.com/AlexeyNikitin01/lms-user/internal/ports/httpgin"
)

type TestServer struct {
	Client  *http.Client
	BaseURL string
}

func CreateTestServer() TestServer {
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

	domainUser := app.CreateAppUser(postgres.CreateRepoUser(db), nil)

	svr := httpgin.Server(":18080", domainUser)

	testServer := httptest.NewServer(svr.Handler)

	return TestServer{
		Client:  testServer.Client(),
		BaseURL: testServer.URL,
	}
}

func (tc *TestServer) GetResponse(req *http.Request, out any) error {
	resp, err := tc.Client.Do(req)
	if err != nil {
		return fmt.Errorf("unexpected error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response: %w", err)
	}

	err = json.Unmarshal(respBody, out)
	if err != nil {
		return fmt.Errorf("unable to unmarshal: %w", err)
	}

	return nil
}

func CreateTestGRPC() {
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

	domainUser := app.CreateAppUser(postgres.CreateRepoUser(db), nil)

	svc := grpcPort.NewService(
		domainUser,
	)

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(svc.Interceptor()))

	grpcPort.RegisterUserServiceServer(grpcServer, svc)

	port := ":50054"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	_ = grpcServer.Serve(lis)

}

func GetClientGRPC() (grpcPort.UserServiceClient, context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	lis := bufconn.Listen(1024 * 1024)

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

	domainUser := app.CreateAppUser(postgres.CreateRepoUser(db), nil)

	svc := grpcPort.NewService(
		domainUser,
	)

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(svc.Interceptor()))
	defer grpcServer.Stop()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	conn, _ := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	defer conn.Close()

	client := grpcPort.NewUserServiceClient(conn)

	return client, ctx
}

func Client(t testing.TB) (grpcPort.UserServiceClient, context.Context) {
	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

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

	domainUser := app.CreateAppUser(postgres.CreateRepoUser(db), nil)

	svc := grpcPort.NewService(
		domainUser,
	)

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(svc.Interceptor()))

	t.Cleanup(func() {
		grpcServer.Stop()
	})

	grpcPort.RegisterUserServiceServer(grpcServer, svc)

	go func() {
		assert.NoError(t, grpcServer.Serve(lis), "srv.Serve")
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	assert.NoError(t, err, "grpc.DialContext")

	t.Cleanup(func() {
		conn.Close()
	})

	return grpcPort.NewUserServiceClient(conn), ctx
}
