package app

import (
	"context"
	"github.com/pkg/errors"
	"lms-user/internal/repository/pg/entity"
)

func (a appUser) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	updatedUser, err := a.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update user")
	}

	return updatedUser, nil
}
