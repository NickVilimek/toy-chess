package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient is a global variable representing the MongoDB client.
var MongoClient *mongo.Client

// ConnectToDB initializes a connection to MongoDB.
func InitMongoConnection(connectionString string) {
	log.Println("Connecting to Mongo")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	MongoClient = client
	log.Println("Connected to MongoDB")
}
