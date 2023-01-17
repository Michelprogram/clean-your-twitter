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

	fmt.Println(tweets_ids)

	var remove int = 50

	/*
		Nombre à supprimer de base à 50 voir 49 pour éviter l'erreur
		Si plus de 50 tweets alors on cap à 50 sinon, on fais une boucle du nombre de tweet

		Lance soit 50 goroutines ou len(tweets)

		- Aucune problème l'utilisateur à ses 50 requests pour delete
		- Problème l'utilisateur manque de ses 50 requests
			- Cancel les goroutines
			- Ajouts les ids à supprimer en bdd

		50 tweets remove per 15 minutes

		Donc si 384 tweets à remove
		384/50 = 7.68
		15 * 7 = 105 minutes
		105/60 = 1.75
		0.75 * 60 = 45
		Donc 1h45

		Pour le script :
		 - Toutes les 15 minutes
		 - Regardes chaque documents de la collection
		 - Supprime les 50 premiers ou moins si utilisé entre temps
		 - Additionne le nombre de tweet delete au fur à mesure
		 - Si arrive à 0 supprime le document et envoie un auto dm twitter comme quoi c'est finito

	*/

	//S'il y a plus de 50 tweets à supprimer
	if len(tweets_ids.TweetsIDS) > 50 {

		//Prend les 50 premiers
		/* 		for _, tweets_id := range tweets_ids.TweetsIDS[50:] {
			//go twitter_api.RemoveTweets(tweets_id)
		} */

		//S'ils en restent stocker en base les suivants
		_, err = dao.AddTweets(tweets_ids.TweetsIDS[50:], *twitter_api.Token, twitter_id)

		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
			return
		}

	} else {
		for _, tweets_id := range tweets_ids.TweetsIDS {
			fmt.Println(tweets_id)
			//go twitter_api.RemoveTweets(tweets_id)
		}

		remove = len(tweets_ids.TweetsIDS)
	}

	text := fmt.Sprintf("{tweet_remove: %d}", remove)

	data, _ := json.Marshal(text)

	fmt.Fprintf(w, "%s", data)
}
