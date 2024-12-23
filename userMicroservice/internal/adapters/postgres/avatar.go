package postgres

import (
	"context"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"lms-user/internal/repository/pg/entity"
)

// SaveAvatarUrl TODO: необходимо проверять наличие аватара.
func (r *RepoUser) SaveAvatarUrl(ctx context.Context, url, userID string) error {
	_, err := entity.Users(entity.UserWhere.ID.EQ(userID)).UpdateAll(ctx, boil.GetContextDB(), entity.M{
		"url": url,
	})
	if err != nil {
		return errors.Wrap(err, "don`t insert db url")
	}

	return nil
}
