package app

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"course/internal/repository/pg/entity"
)

func (c CourseApp) FindLecture(ctx context.Context, courseID int64) (*bson.M, error) {
	lecture, err := c.Mongo.FindLecture(ctx, int(courseID))
	if err != nil {
		return nil, err
	}

	return lecture, nil
}

func (c CourseApp) AllCourse(ctx context.Context, limit, offset int64) (entity.CourseSlice, int64, error) {
	courses, total, err := c.DB.AllCourse(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return courses, total, nil
}

func (c CourseApp) GetCourse(ctx context.Context, courseID int64) (*entity.Course, error) {
	course, err := c.DB.GetCourse(ctx, courseID)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (c CourseApp) GetListUserCourseByID(ctx context.Context, courseID int64) (entity.UsersCourseSlice, error) {
	return c.DB.GetListUserByCourseID(ctx, courseID)
}

func (c CourseApp) GetUserRole(ctx context.Context, courseID int64, uuid string) (*entity.UsersCourse, error) {
	return c.DB.GetUserRole(ctx, courseID, uuid)
}

func (c CourseApp) GetListLectures(ctx context.Context) (entity.LectureSlice, error) {
	return nil, nil
}
