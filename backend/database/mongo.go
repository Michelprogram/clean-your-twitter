package database

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ClientDB *Client

var database string = "clean-your-twitter"

var collection string = "users"

type Client struct {
	Collection *mongo.Collection
	Database   *mongo.Database
}

// Open Connection with database
func OpenConnection() error {

	var err error

	ClientDB, err = Connect()

	if err != nil {
		return err
	}

	return nil
}

// Add Database and collection to the client
func Connect() (*Client, error) {

	uri := os.Getenv("URI_MONGO")

	if uri == "" {
		return nil, errors.New("there is no URI_MONGO environnement variable")
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	collection := client.Database(database).Collection(collection)

	return &Client{
		Collection: collection,
		Database:   client.Database(database),
	}, nil

}
