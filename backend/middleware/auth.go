package middleware

import (
	"api-clean-twitter/jwt"
	"fmt"
	"net/http"
)

func AuthBackend(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		twitter_id, err := r.Cookie("token-twitter")

		if err != nil {
			fmt.Fprintf(w, "no cookie")
			return
		}

		//Decode jwt to get twitter id
		twitter_id.Value, err = jwt.DecodeJwt(twitter_id.Value)

		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
			return
		}

		w.Header().Add("twitter_id", twitter_id.Value)

		next.ServeHTTP(w, r)
	})
}
