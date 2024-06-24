package postgres

import (
	"edu-material/userMicroservice/internal/repository/pg/entity"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/net/context"
)

func (r *RepoUser) Register(ctx context.Context, login, password string) (*entity.User, error) {
	u := &entity.User{
		Login:    login,
		Password: []byte(password),
	}

	err := u.Insert(ctx, r.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	err = u.Reload(ctx, boil.GetContextDB())
	if err != nil {
		return nil, err
	}

	return u, nil
}
