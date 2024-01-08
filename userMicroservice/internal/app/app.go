package app

import (
	"context"
	"edu-material/userMicroservice/internal/user"
)

type AppUser interface {
	CreateUser(ctx context.Context, user user.User) (*user.User, error)
	DeleteUser(ctx context.Context, id int) (*user.User, error)
	GetUser(ctx context.Context, id int) (*user.User, error)
	UpdateUser(ctx context.Context, user user.User) (*user.User, error)
}

type appUser struct {
	repo user.RepoUser
}

func(a appUser) CreateUser(ctx context.Context, user user.User) (*user.User, error) {
	id, err := a.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return &user, nil
}

func(a appUser) DeleteUser(ctx context.Context, id int) (*user.User, error) {return nil, nil} 
func(a appUser) GetUser(ctx context.Context, id int) (*user.User, error) {return nil, nil}
func(a appUser) UpdateUser(ctx context.Context, user user.User) (*user.User, error) {return nil, nil}

func CreateAppUser(repoUser user.RepoUser) AppUser {
	return &appUser{
		repo: repoUser,
	}
}
