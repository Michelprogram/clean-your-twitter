package twitter

import (
	"errors"
	"fmt"
	"os"
)

type DataUser struct {
	Data struct {
		ProfileImageURL string `json:"profile_image_url,omitempty" bson:"profile_image_url"`
		Username        string `json:"username,omitempty" bson:"username"`
		Name            string `json:"name,omitempty" bson:"name"`
		PublicMetrics   struct {
			FollowersCount int `json:"followers_count,omitempty" bson:"followers_count"`
			FollowingCount int `json:"following_count,omitempty" bson:"following_count"`
			TweetCount     int `json:"tweet_count,omitempty" bson:"tweet_count"`
		} `json:"public_metrics,omitempty" bson:"public_metrics"`
		TwitterId string `json:"id,omitempty" bson:"twitter_id"`
	} `json:"data,omitempty"`
}

func NewDataUser(imageUrl, username, twitter_id, name string, followersCount, followingCount int) *DataUser {
	u := &DataUser{}

	u.Data.ProfileImageURL = imageUrl
	u.Data.Username = username
	u.Data.TwitterId = twitter_id
	u.Data.Name = name
	u.Data.PublicMetrics.FollowersCount = followersCount
	u.Data.PublicMetrics.FollowingCount = followingCount
	u.Data.PublicMetrics.TweetCount = 0

	return u
}

func (u DataUser) String() string {
	return fmt.Sprintf("Name : %s\nUsername : %s\nTwitter id : %s\nPicture : %s\nTweet count : %d\nFollower count : %d\nFollowing count :%d\n",
		u.Data.Name, u.Data.Username, u.Data.TwitterId, u.Data.ProfileImageURL, u.Data.PublicMetrics.TweetCount, u.Data.PublicMetrics.FollowersCount, u.Data.PublicMetrics.FollowingCount)
}

type Token struct {
	Expires int16  `json:"expires_in,omitempty" bson:"expires_in"`
	Access  string `json:"access_token,omitempty" bson:"access_token"`
	Refresh string `json:"refresh_token,omitempty" bson:"refresh_token"`
}

func (t Token) String() string {
	return fmt.Sprintf("Token : %s\nRefresh : %s\nExpire in : %d\n", t.Access, t.Refresh, t.Expires)
}

func NewToken(expires int16, access, refresh string) *Token {

	return &Token{
		Expires: expires,
		Access:  access,
		Refresh: refresh,
	}

}

type Twitter struct {
	Client       string `json:"client_id"`
	SecretClient string `json:"secret_client_id"`
	RedirectUri  string `json:"redirect_uri"`

	Token *Token
}

func NewTwitter() (*Twitter, error) {

	client_id := os.Getenv("CLIENT_ID")
	if client_id == "" {
		return nil, errors.New("there is no CLIENT_ID environnement variable")
	}

	secret_client_id := os.Getenv("SECRET_CLIENT_ID")
	if secret_client_id == "" {
		return nil, errors.New("there is no SECRET_CLIENT_ID environnement variable")
	}

	redirect_uri := os.Getenv("REDIRECT_URI")
	if redirect_uri == "" {
		return nil, errors.New("there is no REDIRECT_URI environnement variable")
	}

	return &Twitter{
		Client:       client_id,
		SecretClient: secret_client_id,
		RedirectUri:  redirect_uri,

		Token: &Token{},
	}, nil

}

func (t *Twitter) SetToken(token *Token) {
	t.Token = token
}

func (t Twitter) String() string {
	return fmt.Sprintf("Client id : %s\nSecret id : %s\nRedirect uri : %s\n%s", t.Client, t.SecretClient, t.RedirectUri, t.Token)
}
