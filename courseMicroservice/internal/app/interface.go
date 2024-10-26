package app

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"course/internal/repository/pg/entity"
)

type ICourseApp interface {
	AddCourse(ctx context.Context, name string, description string) (*entity.Course, error)
	AddLecture(ctx context.Context, title, lecture string, courseID int) error
	FindLecture(ctx context.Context, courseID int) (*bson.M, error)
	AllCourse(ctx context.Context, limit, offset int64) (entity.CourseSlice, error)
}
