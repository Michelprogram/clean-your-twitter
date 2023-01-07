package database_test

import (
	"api-clean-twitter/database"
	"context"
	"testing"
)

var ctx = context.TODO()

var client *database.Client

var err error

func TestClient(t *testing.T) {

	client, err = database.SetClient()

	if err != nil {
		t.Errorf("Error during SetClient : got %s want %s\n", err, "nil")
	}

}

func TestConnection(t *testing.T) {

	mongo_client, err := database.Connect(client)

	if err != nil {
		t.Errorf("Error during Connect : got %s want %s\n", err, "nil")
	}

	err = mongo_client.Ping(ctx, nil)

	if err != nil {
		t.Errorf("Error during Ping : got %s want %s\n", err, "nil")
	}

}
