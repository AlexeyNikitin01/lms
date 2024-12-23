package postgres

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"lms-user/internal/repository/pg/entity"
)

func (r *RepoUser) GetUserByTokenDB(ctx context.Context, tokenID string) (*entity.User, *entity.Token, error) {
	tokenUser, err := entity.Tokens(
		entity.TokenWhere.ID.EQ(tokenID),
		entity.TokenWhere.ExpiresAt.GTE(time.Now()),
	).One(ctx, boil.GetContextDB())
	if err != nil {
		return nil, nil, err
	}

	user, err := entity.Users(entity.UserWhere.ID.EQ(tokenUser.UserID)).One(ctx, boil.GetContextDB())
	if err != nil {
		return nil, nil, err
	}

	return user, tokenUser, nil
}

func (r *RepoUser) InsertTokenDB(ctx context.Context, token *entity.Token) error {
	err := token.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return errors.Wrap(err, "AuthByLoginPassword - failed to insert token")
	}

	return nil
}

func (r *RepoUser) GetTokenByUserID(ctx context.Context, userID string) (*entity.Token, error) {
	token, err := entity.Tokens(
		entity.TokenWhere.UserID.EQ(userID),
		entity.TokenWhere.ExpiresAt.GTE(time.Now()),
	).One(ctx, boil.GetContextDB())
	if err != nil {
		return nil, err
	}

	return token, nil
}
