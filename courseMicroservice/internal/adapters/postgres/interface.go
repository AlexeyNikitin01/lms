package postgres

import (
	"github.com/jmoiron/sqlx"
)

type ICoursePostgres interface {
}

/*

	Для использования sql-boiler необходимо использовать драйвер DB
	или boil.GetContextDB() на выбор

*/

type RepoUser struct {
	DB *sqlx.DB
}

func CreateRepoUser(db *sqlx.DB) ICoursePostgres {
	return &RepoUser{
		DB: db,
	}
}
