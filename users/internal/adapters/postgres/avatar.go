package postgres

import (
	"context"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"lms-user/internal/repository/pg/entity"
)

func (r *RepoUser) SaveAvatarFileName(ctx context.Context, avatar, userID string) error {
	_, err := entity.Users(entity.UserWhere.ID.EQ(userID)).UpdateAll(ctx, boil.GetContextDB(), entity.M{
		entity.UserColumns.Avatar: avatar,
	})
	if err != nil {
		return errors.Wrap(err, "don`t insert db url")
	}

	return nil
}
