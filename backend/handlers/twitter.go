package handlers

import (
	"api-clean-twitter/database/repository"
	"api-clean-twitter/entities"
	"api-clean-twitter/jwt"
	"api-clean-twitter/twitter"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var client_uri = "http://localhost:3000/"

func AuthentificationTwitter(w http.ResponseWriter, r *http.Request) {

	var err error

	//Code Generate after logging
	code := r.URL.Query().Get("code")

	//Twitter api golang
	twitter, _ := twitter.NewTwitter()

	//Generate token from new code
	err = twitter.GenerateToken(code)

	if err != nil {
		http.Redirect(w, r, client_uri, http.StatusMovedPermanently)
		return
	}

	//Get Info for user
	dataUser, err := twitter.UsersInfo()

	if err != nil {
		http.Redirect(w, r, client_uri, http.StatusMovedPermanently)
		return
	}

	user := entities.NewUser(dataUser, twitter.Token)

	//Add or update user depend on twitter id
	repository.AddUser(*user)

	//Generate new JWT
	jwt, err := jwt.GenerateJWT(user.Profile.Data.TwitterId)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//Creation of uri
	link_uri := fmt.Sprintf("%s?info=%s", client_uri, jwt.Token)

	//Redirect to the application
	http.Redirect(w, r, link_uri, http.StatusSeeOther)
}

func AuthentificationBackend(w http.ResponseWriter, r *http.Request) {

	twitter_id := w.Header().Get("twitter_id")

	//Looking for user in database
	user, err := repository.GetImageUserById(twitter_id)

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

	//Return image profile
	fmt.Fprintf(w, "%s", data)
}

func Find10LastTweet(w http.ResponseWriter, r *http.Request) {

	twitter_id := w.Header().Get("twitter_id")

	//Looking for user in database
	user, err := repository.GetUserByTwitterId(twitter_id)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	twitter_api, _ := twitter.NewTwitter()

	twitter_api.SetToken(user.User.Token)

	res, err := twitter_api.GetTweets(10, twitter_id)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	fmt.Fprintf(w, "%s", res)
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
