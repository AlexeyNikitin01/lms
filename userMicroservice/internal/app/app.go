package app

import (
	"context"
	"edu-material/userMicroservice/internal/adapters/postgres"
	"edu-material/userMicroservice/internal/repository/pg/entity"
)

type IAppUser interface {
	RegisterDB(ctx context.Context, login, password string) (*entity.User, error)
	GetUserDB(ctx context.Context, uuid string) (*entity.User, error)
}

type appUser struct {
	repo postgres.IUserPostgres
}

func CreateAppUser(repoUser postgres.IUserPostgres) IAppUser {
	return &appUser{
		repo: repoUser,
	}
}
