package postgres

import (
	"context"
	"edu-material/userMicroservice/internal/repository/pg/entity"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (r *RepoUser) GetUserByTokenDB(ctx context.Context, tokenID string) (*entity.User, *entity.Token, error) {
	tokenUser, err := entity.Tokens(entity.TokenWhere.ID.EQ(tokenID)).One(ctx, boil.GetContextDB())
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
	//todo: происходит задубливание в бд, для одного пользователя может быть много токенов
	err := token.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return errors.Wrap(err, "AuthByLoginPassword - failed to insert token")
	}

	return nil
}
