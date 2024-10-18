package app

import (
	"context"

	"course/internal/repository/pg/entity"
)

type ICourseApp interface {
	AddCourse(ctx context.Context, name string) (*entity.Course, error)
	AddLecture(ctx context.Context, title, lecture string, courseID int) error
}
