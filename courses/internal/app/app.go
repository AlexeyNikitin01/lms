package app

import (
	"course/internal/adapters/cloud"
	nosql "course/internal/adapters/mongo"
	"course/internal/adapters/postgres"
)

type CourseApp struct {
	DB    postgres.ICoursePostgres
	Mongo nosql.ICourseMongo
	S3 cloud.ICloud
}

func NewCourseApp(db postgres.ICoursePostgres, mongo nosql.ICourseMongo, s3 cloud.ICloud) *CourseApp {
	return &CourseApp{
		DB:    db,
		Mongo: mongo,
		S3: s3,
	}
}
