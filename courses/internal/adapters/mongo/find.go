package nosql

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r RepoMongoCourse) FindLecture(ctx context.Context, courseID int) (*bson.M, error) {
	database := r.Mongo.Database("course")
	collection := database.Collection("lecture")

	var result *bson.M // Для хранения результата поиска

	filter := bson.D{{"courseID", courseID}} // Условие поиска

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
