package handlers

import (
	mongo "api-clean-twitter/database"
	"api-clean-twitter/entities"
	"api-clean-twitter/jwt"
	"api-clean-twitter/twitter"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var client_id string = "UGVlUFB0N0pxeXA2UWhxX0tiZlI6MTpjaQ"

var secrect_client_id string = "QFuVQ7JTCD5q9fbEyAHLdCTfJkfub6WYunGq8z9v4HtCORjEu7"

var redirect_uri string = "http://www.localhost:3021/auth/twitter"

var client_uri = "http://localhost:3000/"

func AuthentificationTwitter(w http.ResponseWriter, r *http.Request) {

	var err error

	code := r.URL.Query().Get("code")

	twitter := twitter.NewTwitter(client_id, secrect_client_id, redirect_uri)

	err = twitter.GenerateToken(code)

	if err != nil {
		http.Redirect(w, r, client_uri, http.StatusMovedPermanently)
	}

	dataUser, err := twitter.UsersInfo()

	if err != nil {
		http.Redirect(w, r, client_uri, http.StatusMovedPermanently)
	}

	user := entities.NewUser(dataUser, twitter.Token)

	mongo.AddUser(*user)

	jwt, err := jwt.GenerateJWT(user.Profile.Data.TwitterId)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(jwt)

	link_uri := fmt.Sprintf("%s?info=%s", client_uri, jwt.Token)

	http.Redirect(w, r, link_uri, http.StatusSeeOther)
}

func AuthentificationBackend(w http.ResponseWriter, r *http.Request) {

	twitter_id := w.Header().Get("twitter_id")

	//Looking for user in database
	user, err := mongo.GetImageUserById(twitter_id)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	//Encode in json
	data, err := json.Marshal(user)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	fmt.Fprintf(w, "%s", data)
}

func Find10LastTweet(w http.ResponseWriter, r *http.Request) {

	twitter_id := w.Header().Get("twitter_id")

	//Looking for user in database
	user, err := mongo.GetUserByTwitterId(twitter_id)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	fmt.Println(user)

	fmt.Fprintf(w, "Test")
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test")
}
