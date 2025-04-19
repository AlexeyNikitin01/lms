package app

import (
	"context"

	"course/internal/repository/pg/entity"
)

func (c CourseApp) AddCourse(ctx context.Context, name string, description string) (*entity.Course, error) {
	course, err := c.DB.AddCourse(ctx, name, description)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (c CourseApp) AddLecture(ctx context.Context, title, lecture string, courseID int64) error {
	err := c.Mongo.AddLecture(ctx, title, lecture, int(courseID))
	if err != nil {
		return err
	}

	return nil
}
