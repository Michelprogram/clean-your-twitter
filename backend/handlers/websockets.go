package handlers

import (
	"api-clean-twitter/dao"
	"api-clean-twitter/jwt"
	"api-clean-twitter/models"
	"api-clean-twitter/twitter"
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"
)

// TODO : ajouter la date de création du compte pour pas research avant

func Fetch(ws *websocket.Conn) {

	var err error
	var next_token string = ""

	defer ws.Close()

	parameters := ws.Request().URL.Query()

	twitter_id := parameters.Get("token-twitter")
	dates := models.Dates{
		Start: parameters.Get("start"),
		End:   parameters.Get("end"),
	}

	if twitter_id == "" || dates.Start == "" || dates.End == "" {
		fmt.Fprintf(ws, "no url parameter")
		return
	}

	//Decode jwt to get twitter id
	twitter_id, err = jwt.DecodeJwt(twitter_id)

	if err != nil {
		fmt.Fprintf(ws, "%s", err.Error())
		return
	}

	//Looking for user in database
	user, err := dao.GetUserByTwitterId(twitter_id)

	if err != nil {
		fmt.Fprintf(ws, "%s", err.Error())
		return
	}

	twitter_api, _ := twitter.NewTwitter()

	twitter_api.SetToken(user.Token)

	//Loop while next token is different ""
	for {
		data, err := twitter_api.GetTweetsBetweenDates(dates, twitter_id, &next_token)

		if err != nil {
			fmt.Fprintf(ws, "%s", err.Error())
			return
		}

		res, _ := json.Marshal(data)

		fmt.Fprintf(ws, "%s", res)

		if next_token == "" {
			break
		}
	}
}

func Deleting(ws *websocket.Conn) {

}
