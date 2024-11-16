package app

import (
	"context"

	"course/internal/repository/pg/entity"
)

func (c CourseApp) AddCourse(ctx context.Context, name string, description string, url string) (*entity.Course, error) {
	course, err := c.DB.AddCourse(ctx, name, description, url)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (c CourseApp) AddLecture(ctx context.Context, title, lecture string, courseID int) error {
	err := c.Mongo.AddLecture(ctx, title, lecture, courseID)
	if err != nil {
		return err
	}

	return nil
}
