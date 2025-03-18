package app

import (
	"context"

	"lms-user/internal/repository/pg/entity"
)

func (a appUser) Register(ctx context.Context, login, password, email string) (*entity.User, error) {
	u, err := a.repo.RegisterDB(ctx, login, password, email)
	if err != nil {
		return nil, err
	}

	a.telemetry.IncSingUp(ctx)

	return u, err
}
