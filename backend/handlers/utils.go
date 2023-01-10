package handlers

import (
	"api-clean-twitter/middleware"

	"errors"
	"fmt"
	"net/http"
	"os"
)

func Test(w http.ResponseWriter, r *http.Request) {

	client_uri := os.Getenv("CLIENT_URI")

	middleware.Write(r, errors.New("Test pong"))

	fmt.Fprintf(w, "Pong : client_uri : %s\n", client_uri)
}

func TestCode(w http.ResponseWriter, r *http.Request) {
	twitter_id := w.Header().Get("twitter_id")

	fmt.Fprintln(w, twitter_id)
}
