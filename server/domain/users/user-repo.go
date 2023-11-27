package users

import (
	"context"
	"errors"
	"log"
	"toy-chess/infra/database"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetMongoUsers(filters *map[string]interface{}) (*[]User, error) {

	log.Println("[user-repo.GetMongoUsers] Requesting users from Mongo")

	collection := database.MongoClient.Database("toy_chess").Collection("users")

	var users []User

	cursor, err := collection.Find(context.Background(), filters)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No Users Found")
			return &users, nil
		}
		log.Printf("Error - %s\n", err)
		return nil, err
	}

	log.Println("Got users, parsing cursor")

	for cursor.Next(context.Background()) {
		log.Println("Looping on cursor")
		var user User
		if err := cursor.Decode(&user); err != nil {
			return &users, errors.New("Cannot parse user here")
		}
		log.Println("Appending user to output")
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error w curors")
	}

	log.Println("Returning")

	return &users, nil

}
