package postgres

import (
	"context"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/AlexeyNikitin01/lms-user/internal/repository/pg/entity"
)

func (r *RepoUser) DeleteOldToken(ctx context.Context, userID string) error {
	_, err := entity.Tokens(
		entity.TokenWhere.UserID.EQ(userID),
	).DeleteAll(ctx, boil.GetContextDB(), true)
	if err != nil {
		return errors.Wrap(err, "DeleteOldToken")
	}

	return nil
}

func (r *RepoUser) DeleteById(ctx context.Context, userID string) error {
	_, err := entity.Users(
		entity.UserWhere.ID.EQ(userID)).DeleteAll(ctx, boil.GetContextDB(), true)
	if err != nil {
		return errors.Wrap(err, "DeleteById")
	}

	return nil
}
