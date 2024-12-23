package app

import (
	"context"
	"mime/multipart"

	"lms-user/internal/adapters/cloud"
	"lms-user/internal/adapters/postgres"
	"lms-user/internal/repository/pg/entity"
)

type IAppUser interface {
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	Register(ctx context.Context, login, password string, email string) (*entity.User, error)
	GetUser(ctx context.Context, uuid string) (*entity.User, error)
	RefreshToken(tokenUser *entity.Token) (string, error)
	AccessToken(user *entity.User, tokenUser *entity.Token) (string, error)
	ParseToken(ctx context.Context, token string) (*entity.User, *entity.Token, error)
	AuthByLoginPassword(ctx context.Context, login, password string) (*entity.User, *entity.Token, error)
	UploadAvatarS3(
		ctx context.Context,
		fileForm multipart.File,
		header *multipart.FileHeader,
		userID string,
	) (string, error)
	GetUsers(ctx context.Context) (entity.UserSlice, error)
}

type appUser struct {
	repo postgres.IUserPostgres
	s3   cloud.ICloud
}

func CreateAppUser(repoUser postgres.IUserPostgres, s3 cloud.ICloud) IAppUser {
	return &appUser{
		repo: repoUser,
		s3:   s3,
	}
}
