package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"course/internal/repository/pg/entity"
)

type ICoursePostgres interface {
	AddCourse(ctx context.Context, name string) (*entity.Course, error)
}

/*

	Для использования sql-boiler необходимо использовать драйвер DB
	или boil.GetContextDB() на выбор

*/

type RepoCourse struct {
	DB *sqlx.DB
}

func CreateRepoUser(db *sqlx.DB) ICoursePostgres {
	return &RepoCourse{
		DB: db,
	}
}
