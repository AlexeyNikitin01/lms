package app

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"course/internal/repository/pg"
	"course/internal/repository/pg/entity"
)

func (c CourseApp) AddCourse(
	ctx context.Context,
	name string,
	description string,
	authorUUID string,
) (course *entity.Course, err error) {
	if err = pg.ExecTx(ctx, func(ctx context.Context, tx boil.ContextExecutor) error {
		course, err = c.DB.AddCourse(ctx, name, description)
		if err != nil {
			return err
		}

		if err = c.DB.SetAuthorCourse(ctx, course.ID, authorUUID); err != nil {
			return err
		}

		course, err = c.DB.GetCourse(ctx, course.ID)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
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
