package app

import (
	"context"

	"github.com/AlexeyNikitin01/lms-user/internal/repository/pg/entity"
)

func (a appUser) GetUser(ctx context.Context, uuid string) (*entity.User, error) {
	u, err := a.repo.GetUserDB(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (a appUser) GetUsers(ctx context.Context) (entity.UserSlice, error) {
	users, err := a.repo.GetUsersDB(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (a appUser) GetUserByLoginPassword(ctx context.Context, login string, password string) (*entity.User, error) {
	u, err := a.repo.GetUserByLoginPasswordDB(ctx, login, password)
	if err != nil {
		return nil, err
	}

	return u, nil
}
