package postgres

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"course/internal/repository/pg/entity"
)

func (r RepoCourse) AddCourse(ctx context.Context, name string, description string, url string) (*entity.Course, error) {
	c := &entity.Course{
		Name:        name,
		Description: description,
		PhotoURL:    url,
	}

	err := c.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return nil, err
	}

	if err := c.Reload(ctx, boil.GetContextDB()); err != nil {
		return nil, err
	}

	return c, nil
}
