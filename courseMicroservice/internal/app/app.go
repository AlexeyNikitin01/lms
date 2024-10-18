package app

import (
	nosql "course/internal/adapters/mongo"
	"course/internal/adapters/postgres"
)

type CourseApp struct {
	DB    postgres.ICoursePostgres
	Mongo nosql.ICourseMongo
}

func NewCourseApp(db postgres.ICoursePostgres, mongo nosql.ICourseMongo) *CourseApp {
	return &CourseApp{
		DB:    db,
		Mongo: mongo,
	}
}
