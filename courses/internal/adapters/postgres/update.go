package postgres

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"course/internal/repository/pg/entity"
)

func (r RepoCourse) UpdateCourse(ctx context.Context, courseID int64, course *entity.Course) error {
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	for _, m := range course.R.GetModules() {
		if err = r.insertModule(ctx, tx, m, courseID); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r RepoCourse) insertModule(
	ctx context.Context,
	tx boil.ContextExecutor,
	module *entity.Module,
	courseID int64,
) error {
	module.CourseID = courseID
	if err := module.Upsert(
		ctx,
		tx,
		true,
		[]string{entity.ModuleColumns.ID},
		boil.Infer(),
		boil.Infer(),
	); err != nil {
		return err
	}

	lectures := module.R.GetLectures()

	if err := module.Reload(ctx, tx); err != nil {
		return err
	}

	for _, l := range lectures {
		if err := r.insertLecture(ctx, tx, module.ID, l); err != nil {
			return err
		}
	}

	return nil
}

func (r RepoCourse) insertLecture(
	ctx context.Context,
	tx boil.ContextExecutor,
	moduleID int64,
	lecture *entity.Lecture,
) error {
	lecture.ModuleID = moduleID

	if err := lecture.Upsert(
		ctx,
		tx,
		true,
		[]string{entity.LectureColumns.ID},
		boil.Infer(),
		boil.Infer()); err != nil {
		return err
	}

	tests := lecture.R.GetTests()

	if err := lecture.Reload(ctx, tx); err != nil {
		return err
	}

	for _, t := range tests {
		if err := r.insertTest(ctx, tx, lecture.ID, t); err != nil {
			return err
		}
	}

	return nil
}

func (r RepoCourse) insertTest(ctx context.Context, tx boil.ContextExecutor, lectureID int64, test *entity.Test) error {
	test.LectureID = lectureID

	if err := test.Upsert(
		ctx,
		tx,
		true,
		[]string{entity.TestColumns.ID},
		boil.Infer(),
		boil.Infer()); err != nil {
		return err
	}

	questions := test.R.GetQuestions()

	if err := test.Reload(ctx, tx); err != nil {
		return err
	}

	for _, q := range questions {
		q.TestID = test.ID

		if err := q.Upsert(ctx,
			tx,
			true,
			[]string{entity.QuestionColumns.ID},
			boil.Infer(),
			boil.Infer()); err != nil {
			return err
		}

		answers := q.R.GetAnswers()

		if err := q.Reload(ctx, tx); err != nil {
			return err
		}

		for _, a := range answers {
			a.QuestionID = q.ID
			if err := a.Upsert(ctx,
				tx,
				true,
				[]string{entity.AnswerColumns.ID},
				boil.Infer(),
				boil.Infer()); err != nil {
				return err
			}
		}
	}

	return nil
}
