package postgres

import (
	"context"

	"lms-user/internal/repository/pg/entity"

	"github.com/pkg/errors"
)

func (r *RepoUser) GetUserDB(ctx context.Context, uuid string) (*entity.User, error) {
	u, err := entity.Users(
		entity.UserWhere.ID.EQ(uuid),
	).One(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *RepoUser) GetUserByLoginPasswordDB(ctx context.Context, login, password string) (*entity.User, error) {
	u, err := entity.Users(
		entity.UserWhere.Login.EQ(login),
		entity.UserWhere.Password.EQ([]byte(password)),
	).One(ctx, r.DB)
	if err != nil {
		return nil, errors.Wrap(err, "err storage get user")
	}

	return u, nil
}
