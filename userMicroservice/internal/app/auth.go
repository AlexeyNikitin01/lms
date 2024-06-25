package app

import (
	"context"
	"edu-material/userMicroservice/internal/repository/pg/entity"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (a appUser) AuthByLoginPassword(ctx context.Context, login, password string) (*entity.User, *entity.Token, error) {
	user, err := a.repo.GetUserByLoginPasswordDB(ctx, login, password)
	if err != nil {
		return nil, nil, errors.Wrap(err, "appUser AuthByLoginPassword")
	}

	tokensUser := &entity.Token{
		ID:     uuid.New().String(),
		UserID: user.ID,
	}

	_, err = a.AccessToken(user, tokensUser)
	if err != nil {
		return nil, nil, errors.Wrap(err, "appUser AuthByLoginPassword")
	}

	_, err = a.RefreshToken(tokensUser)
	if err != nil {
		return nil, nil, errors.Wrap(err, "appUser AuthByLoginPassword")
	}

	err = a.repo.InsertTokenDB(ctx, tokensUser)
	if err != nil {
		return nil, nil, errors.Wrap(err, "appUser AuthByLoginPassword")
	}

	return user, tokensUser, nil
}
