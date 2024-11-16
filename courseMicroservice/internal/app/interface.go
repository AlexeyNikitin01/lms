package app

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"course/internal/repository/pg/entity"
)

type ICourseApp interface {
	AddCourse(ctx context.Context, name string, description string, url string) (*entity.Course, error)
	AddLecture(ctx context.Context, title, lecture string, courseID int) error
	AllCourse(ctx context.Context, limit, offset int64) (entity.CourseSlice, int64, error)
	UploadPhoto(ctx context.Context, photo []byte, filename string, mime string) (string, error)
	FindLecture(ctx context.Context, courseID int) (*bson.M, error)
}
