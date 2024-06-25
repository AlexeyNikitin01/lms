package postgres

import (
	"context"
	"edu-material/userMicroservice/internal/repository/pg/entity"
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
