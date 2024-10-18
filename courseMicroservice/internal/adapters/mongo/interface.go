package nosql

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ICourseMongo interface {
	AddLecture(ctx context.Context, title, lecture string, courseID int) error
}

type RepoMongoCourse struct {
	Mongo *mongo.Client
}

func NewMongoRepo(mongo *mongo.Client) ICourseMongo {
	return &RepoMongoCourse{
		Mongo: mongo,
	}
}
