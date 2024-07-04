package app

import (
	"context"
	"lms-user/internal/adapters/postgres"
	"lms-user/internal/repository/pg/entity"
)

type IAppUser interface {
	Register(ctx context.Context, login, password string) (*entity.User, error)
	GetUser(ctx context.Context, uuid string) (*entity.User, error)
	RefreshToken(tokenUser *entity.Token) (string, error)
	AccessToken(user *entity.User, tokenUser *entity.Token) (string, error)
	ParseToken(ctx context.Context, token string) (*entity.User, *entity.Token, error)
	AuthByLoginPassword(ctx context.Context, login, password string) (*entity.User, *entity.Token, error)
}

type appUser struct {
	repo postgres.IUserPostgres
}

func CreateAppUser(repoUser postgres.IUserPostgres) IAppUser {
	return &appUser{
		repo: repoUser,
	}
}
