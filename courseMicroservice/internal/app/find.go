package app

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (c CourseApp) FindLecture(ctx context.Context, courseID int) (*bson.M, error) {
	lecture, err := c.Mongo.FindLecture(ctx, courseID)
	if err != nil {
		return nil, err
	}

	return lecture, nil
}
