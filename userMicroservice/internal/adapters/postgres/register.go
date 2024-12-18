package postgres

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/net/context"

	"github.com/AlexeyNikitin01/lms-user/internal/repository/pg/entity"
)

func (r *RepoUser) RegisterDB(ctx context.Context, login, password, email string) (*entity.User, error) {
	u := &entity.User{
		Login:    login,
		Password: []byte(password),
		Email:    email,
	}

	err := u.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return nil, err
	}

	err = u.Reload(ctx, boil.GetContextDB())
	if err != nil {
		return nil, err
	}

	return u, nil
}
