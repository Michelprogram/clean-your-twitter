package models

import "fmt"

type User struct {
	ProfilePicture string   `json:"profile_image_url,omitempty" bson:"profile_image_url"`
	Username       string   `json:"username,omitempty" bson:"username"`
	Name           string   `json:"name,omitempty" bson:"name"`
	CreatedAt      string   `json:"created_at,omitempty" bson:"created_at"`
	TwitterId      string   `json:"id,omitempty" bson:"twitter_id"`
	TweetDeleted   int      `json:"tweet_deleted,omitempty" bson:"tweet_deleted"`
	Metrics        *Metrics `json:"public_metrics,omitempty" bson:"public_metrics"`
	Token          *Token   `json:"token,omitempty" bson:"token"`
}

type UserMongo struct {
	User User `json:"user,omitempty" bson:"user"`
}

func NewUser(picture, username, name, created_at, twitter_id string, tweet_deleted int, metrics *Metrics, token *Token) *User {
	return &User{
		ProfilePicture: picture,
		Username:       username,
		Name:           name,
		CreatedAt:      created_at,
		Metrics:        metrics,
		Token:          token,
		TweetDeleted:   tweet_deleted,
		TwitterId:      twitter_id,
	}
}

func (u User) String() string {
	return fmt.Sprintf("Pciture : %s\n", u.ProfilePicture)
}
