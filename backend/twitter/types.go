package twitter

import (
	"fmt"
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

func NewTwitter(client_id, secrect_client_id, redirect_uri string) *Twitter {

	return &Twitter{
		Client:       client_id,
		SecretClient: secrect_client_id,
		RedirectUri:  redirect_uri,

		Token: &Token{},
	}

}

func (t *Twitter) SetToken(token *Token) {
	t.Token = token
}

func (t Twitter) String() string {
	return fmt.Sprintf("Client id : %s\nSecret id : %s\nRedirect uri : %s\n%s", t.Client, t.SecretClient, t.RedirectUri, t.Token)
}
