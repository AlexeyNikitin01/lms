package client

import (
	"context"
	"lms-user/internal/ports/grpc"
	"log"

	connect "google.golang.org/grpc"
)

const address = "localhost:50051"

func NewClientGRPC(ctx context.Context) grpc.UserServiceClient {
	conn, err := connect.DialContext(ctx, address)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("client connected to ", address)

	return grpc.NewUserServiceClient(conn)
}
