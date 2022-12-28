package entities

import (
	"api-clean-twitter/twitter"
)

type User struct {
	Profile      *twitter.DataUser `json:"profile,omitempty" bson:"profile"`
	Token        *twitter.Token    `json:"token,omitempty" bson:"token"`
	TweetDeleted int               `json:"tweet_deleted,omitempty" bson:"tweet_deleted"`
}

func NewUser(profile *twitter.DataUser, token *twitter.Token) *User {
	return &User{
		Profile: profile,
		Token:   token,
	}
}

type UserMongo struct {
	User User `json:"user,omitempty" bson:"user"`
}

