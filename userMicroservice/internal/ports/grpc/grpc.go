package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/lms-user/internal/app"
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
	Interceptor() grpc.UnaryServerInterceptor
}

type gRPCServerStruct struct {
	UnimplementedUserServiceServer // Вызывает панику, если метод не реализован
	domainUser                     app.IAppUser
}

func (s gRPCServerStruct) Ping(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	log.Println("ping server")

	return &emptypb.Empty{}, nil
}

func NewService(domainUser app.IAppUser) IServer {
	return &gRPCServerStruct{
		domainUser: domainUser,
	}
}
