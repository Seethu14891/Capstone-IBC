package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database
var DatabaseName = "cielcanine"
var DatabaseURI = "mongodb://localhost:27017"

func init() {
	// DatabaseURI := "mongodb://localhost:27017"
	// DatabaseName := "cielcanine" // Set the database name here

	clientOptions := options.Client().ApplyURI(DatabaseURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	Database = client.Database(DatabaseName)
}
