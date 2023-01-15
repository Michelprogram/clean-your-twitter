package models

type Tweets struct {
	TweetsIDS []string `json:"tweets_id, omitempty" bson:"tweets_ids"`
	Token     Token    `json:"token,omitempty" bson:"token"`
	TwitterID string   `json:"id,omitempty" bson:"twitter_id"`
}

func NewTweets(tweets_ids []string, token Token) *Tweets {
	return &Tweets{
		TweetsIDS: tweets_ids,
		Token:     token,
	}
}
