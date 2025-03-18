package app

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"lms-user/internal/repository/pg/entity"
)

var point string = "appUser AuthByLoginPassword"

func (a appUser) AuthByLoginPassword(ctx context.Context, login, password string) (*entity.User, *entity.Token, error) {
	user, err := a.repo.GetUserByLoginPasswordDB(ctx, login, password)
	if err != nil {
		return nil, nil, errors.Wrap(err, point)
	}

	oldToken, err := a.repo.GetTokenByUserID(ctx, user.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, nil, errors.Wrap(err, point)
	}

	if oldToken != nil {
		a.telemetry.IncSingIn(ctx)

		return user, oldToken, nil
	}

	tokensUser := &entity.Token{
		ID:     uuid.New().String(),
		UserID: user.ID,
	}

	_, err = a.AccessToken(user, tokensUser)
	if err != nil {
		return nil, nil, errors.Wrap(err, point)
	}

	_, err = a.RefreshToken(tokensUser)
	if err != nil {
		return nil, nil, errors.Wrap(err, point)
	}

	err = a.repo.DeleteOldToken(ctx, user.ID)
	if err != nil {
		return nil, nil, errors.Wrap(err, point)
	}

	err = a.repo.InsertTokenDB(ctx, tokensUser)
	if err != nil {
		return nil, nil, errors.Wrap(err, point)
	}

	a.telemetry.IncSingIn(ctx)

	return user, tokensUser, nil
}
