package postgres

import (
	"context"
	"edu-material/userMicroservice/internal/repository/pg/entity"
	"github.com/jmoiron/sqlx"
)

type IUserPostgres interface {
	RegisterDB(ctx context.Context, login, password string) (*entity.User, error)
	GetUserDB(ctx context.Context, uuid string) (*entity.User, error)
	GetUserByLogin(ctx context.Context, login, password string) (*entity.User, error)
}

/*

	Для использования sql-boiler необходимо использовать драйвер DB
	или boil.GetContextDB() на выбор

*/

type RepoUser struct {
	DB *sqlx.DB
}

func CreateRepoUser(db *sqlx.DB) IUserPostgres {
	return &RepoUser{
		DB: db,
	}
}
