package app

import (
	"context"
	"edu-material/userMicroservice/internal/repository/pg/entity"
)

func (a appUser) GetUserDB(ctx context.Context, uuid string) (*entity.User, error) {
	u, err := a.repo.GetUserDB(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return u, nil
}
