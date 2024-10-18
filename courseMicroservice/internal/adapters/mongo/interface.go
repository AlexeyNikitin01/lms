package nosql

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ICourseMongo interface {
}

type RepoMongoCourse struct {
	Mongo *mongo.Client
}

func NewMongoRepo(mongo *mongo.Client) ICourseMongo {
	return &RepoMongoCourse{
		Mongo: mongo,
	}
}
