package cron

import (
	"api-clean-twitter/dao"
	"api-clean-twitter/models"
	"api-clean-twitter/twitter"
	"fmt"
	"strconv"
	"time"
)

func CleanTweets() {

	for {
		tweets, _ := dao.FetchAllTweets()

		for _, tweetAccount := range tweets {
			go cleanAccount(tweetAccount)
		}

		time.Sleep(time.Minute * 20)
	}

}

func cleanAccount(tweetAccount models.Tweets) {

	var err error
	flag := false

	tweets := tweetAccount.TweetsIDS
	firstTweet := tweets[0]

	//Minus 1 because we remove first tweet every time
	size := len(tweets) - 1

	//Twitter api account
	twitter_api, _ := twitter.NewTwitter()

	twitter_api.SetToken(&tweetAccount.Token)

	//Faire une premiere request de suppression
	rate, err := twitter_api.RemoveTweet(firstTweet)

	if err != nil {
		fmt.Printf("Can't remove first tweet for user %s, err = %s\n", tweetAccount.TwitterID, err)
		return
	}

	//Récupérer le nombre de requets restant stocké dans rate
	remaining, _ := strconv.Atoi(rate.Remaining)

	//Si moins de tweet que de request
	if size < remaining {
		flag = true
		remaining = size
	}

	if size > 0 {
		//Remove remaining tweets
		selected_tweets := tweets[:remaining]

		//Request available
		if remaining != 0 {
			for _, tweet_id := range selected_tweets {
				go twitter_api.RemoveTweet(tweet_id)
			}
		}

		//Remove tweets from database
		err = dao.RemoveTweets(selected_tweets, tweetAccount.TwitterID)

		if err != nil {
			fmt.Printf("Can't remove tweets for user %s, err = %s\n", tweetAccount.TwitterID, err)
			return
		}
	}

	if flag {
		//Send message to twitter
		err = twitter_api.SendMessage(tweetAccount.TwitterID, "Congratulation all your tweets have been deleted 🧙‍♂️")
		if err != nil {
			fmt.Println("Can't send message : ", err)
			return
		}
		//Remove from collection
		err = dao.RemoveUserFromTweets(tweetAccount.TwitterID)
		if err != nil {
			fmt.Println("Can't remove user from tweets collection : ", err)
			return
		}

		fmt.Println("All tweets have been removed")
	}

}
