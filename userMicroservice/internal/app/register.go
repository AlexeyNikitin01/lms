package app

import (
	"context"
	"edu-material/userMicroservice/internal/repository/pg/entity"
)

func (a appUser) RegisterDB(ctx context.Context, login, password string) (*entity.User, error) {
	u, err := a.repo.RegisterDB(ctx, login, password)
	if err != nil {
		return nil, err
	}

	return u, err
}
