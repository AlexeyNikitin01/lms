package nosql

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// AddLecture TODO: вынести в постоянные переменные.
func (r RepoMongoCourse) AddLecture(ctx context.Context, title, lecture string, courseID int) error {
	database := r.Mongo.Database("course")
	collection := database.Collection("lecture")

	doc := bson.D{
		{"title", title},
		{"lecture", lecture},
		{"courseID", courseID},
	}

	_, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	return nil
}
