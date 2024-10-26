package postgres

import (
	"context"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"course/internal/repository/pg/entity"
)

func (r RepoCourse) AllCourse(ctx context.Context, limit, offset int64) (entity.CourseSlice, error) {
	courses, err := entity.Courses(
		qm.Limit(int(limit)),
		qm.Offset(int(offset)),
	).All(ctx, boil.GetContextDB())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all courses")
	}

	return courses, nil
}
