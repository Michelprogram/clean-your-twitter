package twitter

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func createBasicHeader(twitter *Twitter) string {
	concatened := fmt.Sprintf("%s:%s", twitter.Client, twitter.SecretClient)

	return "Basic " + base64.StdEncoding.EncodeToString([]byte(concatened))
}

func createBearerString(token *Token) string {
	return fmt.Sprintf("Bearer %s", token.Access)
}

func getHTTP(url string, twitter *Twitter) (string, error) {

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Authorization", createBearerString(twitter.Token))
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(resp_body), nil
}

func postHTTP(data url.Values, url string, twitter *Twitter) (string, error) {

	encoded := data.Encode()

	resp, _ := http.NewRequest("POST", url, strings.NewReader(encoded))
	resp.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	resp.Header.Set("Authorization", createBasicHeader(twitter))

	client := &http.Client{}
	response, error := client.Do(resp)
	if error != nil {
		return "", error
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	return string(body), nil
}
