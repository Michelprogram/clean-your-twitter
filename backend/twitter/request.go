package twitter

import (
	"api-clean-twitter/dao"
	"api-clean-twitter/models"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	data    url.Values
	url     string
	twitter Twitter
}

func NewRequest(url string, twitter Twitter, data url.Values) *Request {
	return &Request{
		data:    data,
		url:     url,
		twitter: twitter,
	}
}

func (r Request) createBasicHeader() string {
	concatened := fmt.Sprintf("%s:%s", r.twitter.Client, r.twitter.SecretClient)

	return "Basic " + base64.StdEncoding.EncodeToString([]byte(concatened))
}

func (r Request) createBearerString() string {
	return fmt.Sprintf("Bearer %s", r.twitter.Token.Access)
}

func (r *Request) isTokenValid() (bool, error) {

	if !r.twitter.Token.IsValid() && r.url != "https://api.twitter.com/2/oauth2/token" {

		old_token := r.twitter.Token

		err := r.twitter.RefreshToken()

		if err != nil {
			return false, err
		}

		dao.UpdateTokenUser(*r.twitter.Token, old_token.Access)

		return false, err
	}

	return true, nil
}

func (r *Request) GetHTTP() (string, *models.Rate, error) {

	var err error

	_, err = r.isTokenValid()

	if err != nil {
		return "", nil, err
	}

	req, err := http.NewRequest("GET", r.url, nil)

	if err != nil {
		return "", nil, err
	}

	req.Header.Set("Authorization", r.createBearerString())
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", nil, err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", nil, err
	}

	if resp.StatusCode != 200 {
		return "", nil, errors.New(string(resp_body))
	}

	header := &models.Rate{
		Limit:     resp.Header.Get("x-rate-limit-limit"),
		Remaining: resp.Header.Get("x-rate-limit-remaining"),
	}

	return string(resp_body), header, nil
}

func (r *Request) PostHTTP() (string, error) {

	var err error

	_, err = r.isTokenValid()

	if err != nil {
		return "", err
	}

	encoded := r.data.Encode()

	resp, err := http.NewRequest("POST", r.url, strings.NewReader(encoded))

	if err != nil {
		return "", err
	}

	resp.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	resp.Header.Set("Authorization", r.createBasicHeader())

	client := &http.Client{}
	response, err := client.Do(resp)

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		return "", errors.New(string(body))
	}
	return string(body), nil
}

func (r *Request) DeleteHTTP() (string, error) {
	var err error

	_, err = r.isTokenValid()

	if err != nil {
		return "", err
	}

	resp, err := http.NewRequest("DELETE", r.url, nil)

	if err != nil {
		return "", err
	}

	resp.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	resp.Header.Set("Authorization", r.createBearerString())

	client := &http.Client{}
	response, err := client.Do(resp)

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		return "", errors.New(string(body))
	}
	return string(body), nil

}
