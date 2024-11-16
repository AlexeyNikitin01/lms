package postgres

import (
	"context"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/AlexeyNikitin01/lms-user/internal/repository/pg/entity"
)

func (r *RepoUser) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := user.Upsert(ctx, boil.GetContextDB(), true, nil,
		boil.Blacklist("password"), boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "failed to update user storage")
	}

	err = user.Reload(ctx, boil.GetContextDB())
	if err != nil {
		return nil, errors.Wrap(err, "failed to reload user storage")
	}

	return user, nil
}
