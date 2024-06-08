package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type Server interface {
	UserServiceServer
}

type gRPCServerStruct struct {
	UnimplementedUserServiceServer
}

func (s gRPCServerStruct) GetAllUser(ctx context.Context, request *UserRequest) (*emptypb.Empty, error) {
	time.Sleep(time.Microsecond * 100)
	return &emptypb.Empty{}, nil
}

func NewService() Server {
	return &gRPCServerStruct{}
}
