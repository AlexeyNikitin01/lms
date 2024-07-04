package app

import (
	"context"
	"lms-user/internal/repository/pg/entity"
)

func (a appUser) Register(ctx context.Context, login, password string) (*entity.User, error) {
	u, err := a.repo.RegisterDB(ctx, login, password)
	if err != nil {
		return nil, err
	}

	return u, err
}
