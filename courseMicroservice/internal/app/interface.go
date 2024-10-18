package app

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"course/internal/repository/pg/entity"
)

type ICourseApp interface {
	AddCourse(ctx context.Context, name string) (*entity.Course, error)
	AddLecture(ctx context.Context, title, lecture string, courseID int) error
	FindLecture(ctx context.Context, courseID int) (*bson.M, error)
}
