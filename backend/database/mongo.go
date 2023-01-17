package database

import (
	"context"
	"errors"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ClientDB *Client

type Client struct {
	Identifiant string
	Password    string
	URL         string
	Collection  *mongo.Collection
	Database    *mongo.Database
}

// Open Connection with database
func OpenConnection() error {

	var err error

	ClientDB, err = SetClient()

	if err != nil {
		return err
	}

	_, err = Connect(ClientDB)

	if err != nil {
		return err
	}

	return nil
}

// Add Database and collection to the client
func Connect(c *Client) (*mongo.Client, error) {
	//uri := "mongodb://user:password@localhost:27017"
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", c.Identifiant, c.Password, c.URL)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	database := os.Getenv("DATABASE")
	if database == "" {
		return nil, errors.New("there is no DATABASE environnement variable")
	}
	c.Database = client.Database(database)

	collection := os.Getenv("COLLECTION")
	if collection == "" {
		return nil, errors.New("there is no COLLECTION environnement variable")
	}
	c.Collection = c.Database.Collection(collection)

	return client, nil

}

// Set client with env variables
func SetClient() (*Client, error) {

	id := os.Getenv("IDENTIFIANT_DB")
	if id == "" {
		return nil, errors.New("there is no IDENTIFIANT_DB environnement variable")
	}

	password := os.Getenv("PASSWORD_DB")
	if password == "" {
		return nil, errors.New("there is no PASSWORD environnement variable")
	}

	url := os.Getenv("URL_DB")
	if url == "" {
		return nil, errors.New("there is no URL environnement variable")
	}

	return &Client{
		Identifiant: id,
		Password:    password,
		URL:         url,
		Database:    nil,
		Collection:  nil,
	}, nil
}
