package app

import (
	"context"

	"course/internal/repository/pg/entity"
)

type ICourseApp interface {
	AddCourse(ctx context.Context, name string) (*entity.Course, error)
}
