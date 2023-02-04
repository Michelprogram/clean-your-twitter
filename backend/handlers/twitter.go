package handlers

import (
	"api-clean-twitter/dao"
	"api-clean-twitter/jwt"
	"api-clean-twitter/middleware"
	"api-clean-twitter/models"
	"api-clean-twitter/twitter"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func AuthentificationTwitter(w http.ResponseWriter, r *http.Request) {

	var err error

	client_uri := os.Getenv("CLIENT_URI")

	//Code Generate after logging
	code := r.URL.Query().Get("code")

	//Twitter api golang
	twitter, err := twitter.NewTwitter()

	if err != nil {
		middleware.Write(r, err)
		http.Redirect(w, r, client_uri, http.StatusMovedPermanently)
		return
	}

	//Generate token from new code
	err = twitter.GenerateToken(code)

	if err != nil {
		middleware.Write(r, err)
		http.Redirect(w, r, client_uri, http.StatusMovedPermanently)
		return
	}

	//Get Info for user
	data, err := twitter.UsersInfo()

	if err != nil {
		middleware.Write(r, err)
		http.Redirect(w, r, client_uri, http.StatusMovedPermanently)
		return
	}

	metrics := models.NewMetrics(data.Data.PublicMetrics.FollowersCount, data.Data.PublicMetrics.FollowingCount, data.Data.PublicMetrics.TweetCount)

	user := models.NewUser(data.Data.ProfileImageURL, data.Data.Username, data.Data.Name, data.Data.CreatedAt, data.Data.TwitterId, int(0), metrics, twitter.Token)

	//Add or update user depend on twitter id
	dao.AddUser(*user)
	//Update token in tweets collection
	dao.UpdateTokenByTwitterId(*user)

	//Generate new JWT
	jwt, err := jwt.GenerateJWT(user.TwitterId)

	if err != nil {
		middleware.Write(r, err)
		http.Redirect(w, r, client_uri, http.StatusMovedPermanently)
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
	user, err := dao.GetImageUserById(twitter_id)

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

func CleanTweets(w http.ResponseWriter, r *http.Request) {

	var tweets_ids models.CleanTweets
	var stored int

	twitter_id := w.Header().Get("twitter_id")

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tweets_ids); err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	defer r.Body.Close()

	//Looking for user in database
	user, err := dao.GetUserByTwitterId(twitter_id)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	twitter_api, _ := twitter.NewTwitter()

	twitter_api.SetToken(user.Token)

	tweets := tweets_ids.TweetsIDS
	size := len(tweets)

	if size == 0 {
		return
	}

	firstTweet := tweets[0]

	//Faire une premiere request de suppression
	rate, err := twitter_api.RemoveTweet(firstTweet)

	//Récupérer le nombre de requets restant stocké dans rate
	remaining, _ := strconv.Atoi(rate.Remaining)

	//Si moins de tweet que de request
	if size < remaining {
		remaining = size
	}

	//Plus aucune request disponible
	if err != nil || remaining == 0 {
		dao.AddTweets(tweets, *twitter_api.Token, twitter_id)
		stored = size
	} else {
		selected_tweets := tweets[1:remaining]
		for _, tweet_id := range selected_tweets {
			go twitter_api.RemoveTweet(tweet_id)
		}
		//Stocker en base le reste sans doublons dans un tableau
		if size > remaining {
			rest_tweets := tweets[remaining:]
			dao.AddTweets(rest_tweets, *twitter_api.Token, twitter_id)
			stored = len(rest_tweets)
		}
	}

	text := fmt.Sprintf(`{"tweet_removed":"%d", "tweet_stored": "%d"}`, remaining, stored)

	fmt.Fprintf(w, "%s", text)

}
