package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"course/internal/repository/pg/entity"
)

type ICoursePostgres interface {
	AddCourse(ctx context.Context, name string, description string) (*entity.Course, error)
	SetAuthorCourse(ctx context.Context, courseID int64, authorUUID string) error
	AllCourse(ctx context.Context, limit, offset int64) (entity.CourseSlice, int64, error)
	SaveAvatarCourse(ctx context.Context, fileName string, courseID int64) error
	UpdateCourse(ctx context.Context, courseID int64, course *entity.Course) error
	GetCourse(ctx context.Context, courseID int64) (*entity.Course, error)
	GetListUserByCourseID(ctx context.Context, courseID int64) (entity.UsersCourseSlice, error)
	GetUserRole(ctx context.Context, courseID int64, uuid string) (*entity.UsersCourse, error)
	GetListLectures(ctx context.Context) (entity.LectureSlice, error)
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
