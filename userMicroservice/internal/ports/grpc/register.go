package grpc

import (
	"context"
)

/*

lms-1: возможность регистрации пользователей через логин, пароль и email

*/

func (s gRPCServerStruct) RegisterUser(ctx context.Context, req *UserRegisterRequest) (*UserRegisterResponse, error) {
	u, err := s.domainUser.Register(ctx, req.Login, req.Password)
	if err != nil {
		return nil, err
	}

	return &UserRegisterResponse{
		Uuid:  u.ID,
		Login: u.Login,
	}, nil
}
