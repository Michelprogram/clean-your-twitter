package repository_test

import (
	"api-clean-twitter/database"
	"api-clean-twitter/database/repository"
	"api-clean-twitter/entities"
	"api-clean-twitter/twitter"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
)

var err error

var user entities.User = entities.User{
	Profile:      twitter.NewDataUser("test-image", "test-username", "test-twitter-id", "test-name", 0, 0),
	Token:        twitter.NewToken(7200, "test-access", "test-refresh"),
	TweetDeleted: 0,
}

func TestMain(m *testing.M) {

	const PATH string = "../../.env"

	err = godotenv.Load(PATH)
	if err != nil {
		log.Fatal("error loading .env file")
	}

	err = database.OpenConnection()

	if err != nil {
		log.Fatal(err)
	}

	exit_value := m.Run()

	os.Exit(exit_value)
}

func TestAddNewUser(t *testing.T) {
	res, err := repository.AddUser(user)

	if err != nil {
		t.Errorf("Error during AddUser : got %s want %s\n", err, "nil")

	}

	if res.UpsertedCount != 1 {
		t.Errorf("Error during AddUser (upsert) : got %d want %d\n", res.UpsertedCount, 1)
	}

	if res.MatchedCount != 0 {
		t.Errorf("Error during AddUser (matched) : got %d want %d\n", res.MatchedCount, 0)
	}

}

func TestUpdateUser(t *testing.T) {
	res, err := repository.AddUser(user)

	if err != nil {
		t.Errorf("Error during UpdateUser : got %s want %s\n", err, "nil")

	}

	if res.MatchedCount != 1 {
		t.Errorf("Error during UpdateUser (match) : got %d want %d\n", res.MatchedCount, 1)

	}

	if res.UpsertedCount != 0 {
		t.Errorf("Error during UpdateUser (upsert) : got %d want %d\n", res.UpsertedCount, 0)
	}

}

func TestGetUserByTwitterId(t *testing.T) {

	res, err := repository.GetUserByTwitterId("test-twitter-id")

	if err != nil {
		t.Errorf("Error during GetUserByTwitterId : got %s want %s\n", err, "nil")

	}

	got := res.User.Profile.Data.TwitterId
	want := user.Profile.Data.TwitterId

	if got != want {
		t.Errorf("Error during GetUserByTwitterId : got %s want %s\n", got, want)
	}

}

func TestGetImageUserByTwitterId(t *testing.T) {
	res, err := repository.GetImageUserById("test-twitter-id")

	if err != nil {
		t.Errorf("Error during GetImageUserById : got %s want %s\n", err, "nil")

	}

	got := res.User.Profile.Data.ProfileImageURL
	want := user.Profile.Data.ProfileImageURL

	if got != want {
		t.Errorf("Error during GetImageUserById : got %v want %v\n", got, want)
	}
}

func TestUpdateToken(t *testing.T) {

	new_token := twitter.NewToken(7200, "test-access-2", "test-refresh-2")
	twitter_id := user.Profile.Data.TwitterId

	res, err := repository.UpdateTokenUser(twitter_id, new_token)

	if err != nil {
		t.Errorf("Error during UpdateTokenUser : got %s want %s\n", err, "nil")
	}

	if res.ModifiedCount != 1 {
		t.Errorf("Error during UpdateTokenUser (ModifiedCount) : got %d want %d\n", res.MatchedCount, 1)
	}

	u, err := repository.GetUserByTwitterId(twitter_id)

	if err != nil {
		t.Errorf("Error during GetUserByTwitterId : got %s want %s\n", err, "nil")
	}

	if !reflect.DeepEqual(u.User.Token, new_token) {
		t.Errorf("Error during UpdateTokenUser (Token) : got %s want %s\n", u.User.Token, new_token)

	}

}

func TestDeleteUser(t *testing.T) {
	res, err := repository.DeleteUserByTwitterId("test-twitter-id")

	if err != nil {
		t.Errorf("Error during DeleteUser : got %s want %s\n", err, "nil")

	}

	if res.DeletedCount == 0 {
		t.Errorf("Error during DeleteUser : got %d want %d\n", res.DeletedCount, 1)

	}

}
