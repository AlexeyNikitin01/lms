package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"lms-user/internal/repository/pg/entity"
)

type IUserPostgres interface {
	Register(ctx context.Context, login, password string) (*entity.User, error)
	GetUser(ctx context.Context, uuid string) (*entity.User, error)
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
