package app

import (
	"context"

	"course/internal/repository/pg/entity"
)

func (c CourseApp) UpdateCourse(ctx context.Context, courseID int64, course *entity.Course) error {
	err := c.DB.UpdateCourse(ctx, courseID, course)
	if err != nil {
		return err
	}

	return nil
}
