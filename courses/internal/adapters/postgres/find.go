package postgres

import (
	"context"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"course/internal/repository/pg/entity"
)

func (r RepoCourse) AllCourse(ctx context.Context, limit, offset int64) (entity.CourseSlice, int64, error) {
	courses, err := entity.Courses(
		qm.Limit(int(limit)),
		qm.Offset(int(offset)),
	).All(ctx, boil.GetContextDB())
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to get all courses")
	}

	total, err := entity.Courses().Count(ctx, boil.GetContextDB())
	if err != nil {
		return nil, 0, errors.Wrap(err, "count courses")
	}

	return courses, total, nil
}

func (r RepoCourse) GetCourse(ctx context.Context, courseID int64) (*entity.Course, error) {
	course, err := entity.Courses(
		entity.CourseWhere.ID.EQ(courseID),
		qm.Load(qm.Rels(
			entity.CourseRels.Modules,
			entity.ModuleRels.Lectures,
			entity.LectureRels.Tests,
			entity.TestRels.Questions,
			entity.QuestionRels.Answers,
		)),
	).One(ctx, boil.GetContextDB())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get course by id")
	}

	return course, nil
}

func (r RepoCourse) GetListUserByCourseID(ctx context.Context, courseID int64) (entity.UsersCourseSlice, error) {
	listUser, err := entity.UsersCourses(
		entity.UsersCourseWhere.CourseID.EQ(courseID),
	).All(ctx, boil.GetContextDB())
	if err != nil {
		return nil, errors.Wrap(err, "failed to load users courses")
	}

	return listUser, nil
}

func (r RepoCourse) GetUserRole(ctx context.Context, courseID int64, uuid string) (*entity.UsersCourse, error) {
	userRole, err := entity.UsersCourses(
		entity.UsersCourseWhere.CourseID.EQ(courseID),
		entity.UsersCourseWhere.UserUUID.EQ(uuid),
	).One(ctx, boil.GetContextDB())
	if err != nil {
		return nil, errors.Wrap(err, "failed to load users courses")
	}

	return userRole, nil
}
