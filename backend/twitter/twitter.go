package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

var token_url string = "https://api.twitter.com/2/oauth2/token"

func (twitter *Twitter) GenerateToken(code string) error {
	var token Token = Token{}

	data := url.Values{}

	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", twitter.Client)
	data.Set("redirect_uri", twitter.RedirectUri)
	data.Set("code_verifier", "8KxxO-RPl0bLSxX5AWwgdiFbMnry_VOKzFeIlVA7NoA")

	body, err := postHTTP(data, token_url, twitter)

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(body), &token)

	twitter.SetToken(&token)

	return nil
}

func (twitter *Twitter) RefreshToken() error {
	var token Token = Token{}

	data := url.Values{}

	data.Set("refresh_token", twitter.Token.Refresh)
	data.Set("grant_type", "refresh_token")

	body, err := postHTTP(data, token_url, twitter)

	if err != nil {
		return err
	}

	fmt.Println(body)

	json.Unmarshal([]byte(body), &token)

	twitter.SetToken(&token)

	return nil
}

func (twitter Twitter) UsersInfo() (*DataUser, error) {
	var user_info string = "https://api.twitter.com/2/users/me?user.fields=profile_image_url%2Cpublic_metrics"

	var user DataUser

	body, err := getHTTP(user_info, &twitter)

	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(body), &user)

	return &user, nil

}

func (twitter Twitter) GetTweets(n int, twitter_id string) (any, error) {

	if n <= 0 || n > 100 {
		return nil, errors.New("number should be between 1 and 100")
	}

	var path string = fmt.Sprintf("https://api.twitter.com/2/users/%s/tweets?max_results=10&tweet.fields=created_at", twitter_id)
	body, err := getHTTP(path, &twitter)

	if err != nil {
		return nil, err
	}

	return body, nil

}
