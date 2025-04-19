package app

import (
	"context"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/v2/bson"

	"course/internal/repository/pg/entity"
)

type ICourseApp interface {
	AddCourse(ctx context.Context, name string, description string) (*entity.Course, error)
	UpdateCourse(ctx context.Context, courseID int64, course *entity.Course) error
	AddLecture(ctx context.Context, title, lecture string, courseID int64) error
	AllCourse(ctx context.Context, limit, offset int64) (entity.CourseSlice, int64, error)
	UploadPhoto(ctx context.Context, fileForm multipart.File, header *multipart.FileHeader, course *entity.Course) error
	FindLecture(ctx context.Context, courseID int64) (*bson.M, error)
	GetCourse(ctx context.Context, courseID int64) (*entity.Course, error)
}
