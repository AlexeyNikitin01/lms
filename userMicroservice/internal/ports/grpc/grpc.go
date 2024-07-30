package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"lms-user/internal/app"
)

type Server interface {
	UserServiceServer
}

type gRPCServerStruct struct {
	UnimplementedUserServiceServer
	domainUser app.IAppUser
}

func (s gRPCServerStruct) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}

func NewService(domainUser app.IAppUser) Server {
	return &gRPCServerStruct{
		domainUser: domainUser,
	}
}
