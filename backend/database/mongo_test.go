package database_test

import (
	"api-clean-twitter/database"
	"context"
	"testing"
)

func TestConnection(t *testing.T) {

	mongo_client, err := database.Connect()

	if err != nil {
		t.Errorf("Error during Connect : got %s want %s\n", err, "nil")
	}

	err = mongo_client.Database.Client().Ping(context.TODO(), nil)

	if err != nil {
		t.Errorf("Error during Ping : got %s want %s\n", err, "nil")
	}

}
