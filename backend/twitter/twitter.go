package twitter

import (
	"api-clean-twitter/models"
	"encoding/json"
	"fmt"
	"net/url"
)

var token_url string = "https://api.twitter.com/2/oauth2/token"

func (twitter *Twitter) GenerateToken(code string) error {
	var token models.Token = models.Token{}

	data := url.Values{}

	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", twitter.Client)
	data.Set("redirect_uri", twitter.RedirectUri)
	data.Set("code_verifier", "8KxxO-RPl0bLSxX5AWwgdiFbMnry_VOKzFeIlVA7NoA")

	req := NewRequest(token_url, *twitter, data)

	body, err := req.PostHTTP()

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(body), &token)

	token.ConvertToTime()

	twitter.SetToken(&token)

	return nil
}

func (twitter *Twitter) RefreshToken() error {
	var token models.Token = models.Token{}

	data := url.Values{}

	data.Set("refresh_token", twitter.Token.Refresh)
	data.Set("grant_type", "refresh_token")

	req := NewRequest(token_url, *twitter, data)

	body, err := req.PostHTTP()

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(body), &token)

	token.ConvertToTime()

	twitter.SetToken(&token)

	return nil
}

func (twitter Twitter) UsersInfo() (*DataUser, error) {
	var user_info string = "https://api.twitter.com/2/users/me?user.fields=profile_image_url%2Cpublic_metrics"

	var user DataUser

	req := NewRequest(user_info, twitter, nil)

	body, err := req.GetHTTP()

	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(body), &user)

	return &user, nil

}

// TODO : A faire en fonction recursive
func (twitter Twitter) GetTweetsBetweenDates(dates models.Dates, twitter_id string) ([]*models.InfoTweet, error) {

	var flag bool = true
	var url string = ""
	var next_token string = ""
	var data []*models.InfoTweet

	for flag {
		var tweet models.Tweet
		url = fmt.Sprintf("https://api.twitter.com/2/users/%s/tweets?max_results=100&tweet.fields=created_at%%2Centities&exclude=retweets&start_time=%s&end_time=%s", twitter_id, dates.Start, dates.End)

		if next_token != "" {
			url = fmt.Sprintf("%s&=pagination_token=%s", url, next_token)
		}

		req := NewRequest(url, twitter, nil)

		body, err := req.GetHTTP()

		if err != nil {
			return nil, err
		}

		json.Unmarshal([]byte(body), &tweet)

		data = append(data, tweet.Data...)

		if tweet.Meta.NextToken != "" {
			next_token = tweet.Meta.NextToken
		} else {
			flag = false
		}
	}

	return data, nil

}
