package nosql

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongoClient(cfg *MongoConfig) (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.User, cfg.Password, cfg.Host, cfg.Port)),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}
