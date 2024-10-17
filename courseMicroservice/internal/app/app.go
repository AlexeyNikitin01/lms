package app

import (
	"course/internal/adapters/postgres"
)

type CourseApp struct {
	DB postgres.ICoursePostgres
}

func NewCourseApp(db postgres.ICoursePostgres) *CourseApp {
	return &CourseApp{
		DB: db,
	}
}
