package app

import (
	"course/internal/adapters/storage"
	nosql "course/internal/adapters/mongo"
	"course/internal/adapters/postgres"
)

type CourseApp struct {
	DB    postgres.ICoursePostgres
	Mongo nosql.ICourseMongo
	S3    storage.ICloud
}

func NewCourseApp(db postgres.ICoursePostgres, mongo nosql.ICourseMongo, s3 storage.ICloud) *CourseApp {
	return &CourseApp{
		DB:    db,
		Mongo: mongo,
		S3:    s3,
	}
}
