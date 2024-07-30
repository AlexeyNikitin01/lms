package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"lms-user/internal/app"
)

type IServer interface {
	UserServiceServer
}

type gRPCServerStruct struct {
	UnimplementedUserServiceServer // Вызывает панику, если метод не реализован
	domainUser                     app.IAppUser
}

func (s gRPCServerStruct) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}

func NewService(domainUser app.IAppUser) IServer {
	return &gRPCServerStruct{
		domainUser: domainUser,
	}
}
