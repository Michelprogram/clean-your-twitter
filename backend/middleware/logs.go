package middleware

import (
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func init() {
	// open file and create if non-existent
	file, err := os.OpenFile("api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	logger = log.New(file, "", log.LstdFlags)
}

func Logs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt, err := r.Cookie("token-twitter")
		//get ip adresse

		if err != nil {
			logger.Printf(" %s -> %s\n", r.Method, r.URL.Path)
		}

		logger.Printf("%s -> %s jwt : %s\n", r.Method, r.URL.Path, jwt.Value)

		next.ServeHTTP(w, r)
	})
}
