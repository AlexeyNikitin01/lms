package postgres

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"course/internal/repository/pg/entity"
)

func (r RepoCourse) AddCourse(ctx context.Context, name string) (*entity.Course, error) {
	c := &entity.Course{
		Name: name,
	}

	err := c.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return nil, err
	}

	return c, nil
}
