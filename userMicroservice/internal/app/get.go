package app

import (
	"context"
	"lms-user/internal/repository/pg/entity"
)

func (a appUser) GetUser(ctx context.Context, uuid string) (*entity.User, error) {
	u, err := a.repo.GetUser(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return u, nil
}
