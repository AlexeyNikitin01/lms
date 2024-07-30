package grpc

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"lms-user/internal/app"
)

/*

Транспортный слой - отвечает за взаимодействие серверной части приложение с внешним миром.

Основными методы являются:
	- getUser
	- updateUser
	- register
	- authByLoginPassword
	- interceptor for validation jwt token

Авторизация осуществляется через валидацию jwt токена.

*/

type IServer interface {
	UserServiceServer
}

type gRPCServerStruct struct {
	UnimplementedUserServiceServer // Вызывает панику, если метод не реализован
	domainUser                     app.IAppUser
}

func (s gRPCServerStruct) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	log.Println("ping server")
	return &emptypb.Empty{}, nil
}

func NewService(domainUser app.IAppUser) IServer {
	return &gRPCServerStruct{
		domainUser: domainUser,
	}
}
