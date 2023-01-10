package models

type Metrics struct {
	FollowersCount int `json:"followers_count,omitempty" bson:"followers_count"`
	FollowingCount int `json:"following_count,omitempty" bson:"following_count"`
	TweetCount     int `json:"tweet_count,omitempty" bson:"tweet_count"`
}

func NewMetrics(followers_count, following_count, tweet_count int) *Metrics {
	return &Metrics{
		FollowersCount: followers_count,
		FollowingCount: following_count,
		TweetCount:     tweet_count,
	}
}
