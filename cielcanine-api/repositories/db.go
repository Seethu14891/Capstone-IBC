package database

import (
	"cielcanine-api/config" // Import your project's config package
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func Connect() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(config.DatabaseURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	db = client.Database(config.DatabaseName)

	return db, nil
}

func GetDB() *mongo.Database {
	return db
}
