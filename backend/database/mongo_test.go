package database_test

import (
	"api-clean-twitter/database"
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var ctx = context.TODO()

var client *database.Client

var err error

func TestMain(m *testing.M) {

	const PATH string = "../.env"

	err := godotenv.Load(PATH)
	if err != nil {
		log.Fatal("error loading .env file")
	}

	exit_value := m.Run()

	os.Exit(exit_value)
}

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
