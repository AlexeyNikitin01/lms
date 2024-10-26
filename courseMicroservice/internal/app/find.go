package app

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"course/internal/repository/pg/entity"
)

func (c CourseApp) FindLecture(ctx context.Context, courseID int) (*bson.M, error) {
	lecture, err := c.Mongo.FindLecture(ctx, courseID)
	if err != nil {
		return nil, err
	}

	return lecture, nil
}

func (c CourseApp) AllCourse(ctx context.Context, limit, offset int64) (entity.CourseSlice, error) {
	courses, err := c.DB.AllCourse(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return courses, nil
}
