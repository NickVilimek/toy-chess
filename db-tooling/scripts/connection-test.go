package scripts

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        string `bson:"_id,omitempty"`
	FirstName string `bson:"firstname"`
	LastName  string `bson:"lastname"`
	Email     string `bson:"email"`
	// Add other fields as needed
}

func RunLoadUserTest() {

	InitToolingConfig()

	connectionString := DbConfigInstance.MongoConnectionString

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	//Try and get collection value
	ctx := context.Background()

	collection := client.Database("toy_chess").Collection("users")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Maybe got data from collection")

	defer cursor.Close(ctx)

	// Decode the results into a slice of User structs
	var users []User
	for cursor.Next(ctx) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	// Check for errors from iterating over cursor
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Printf("ID: %s, Username: %s, Email: %s\n", user.ID, user.FirstName+" "+user.LastName, user.Email)
	}

}
