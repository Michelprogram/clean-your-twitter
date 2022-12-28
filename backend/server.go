package main

import (
	h "api-clean-twitter/handlers"
	m "api-clean-twitter/middleware"

	"fmt"
	"net/http"
	"os"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func main() {

	var port string = os.Args[1]

	router := mux.NewRouter()

	router.Use(m.Logs)

	//Authentification with twitter
	twitter_routes := router.PathPrefix("/twitter").Subrouter()
	twitter_routes.HandleFunc("/twitter/auth", h.AuthentificationTwitter).Methods("GET")

	//Interaction with backend
	backend_routes := router.PathPrefix("/backend").Subrouter()
	backend_routes.Use(m.AuthBackend)

	backend_routes.HandleFunc("/auth", h.AuthentificationBackend).Methods("GET")
	backend_routes.HandleFunc("/tweets", h.Find10LastTweet).Methods("GET")

	//Ping server
	router.HandleFunc("/test", h.Test).Methods("GET")

	//router.NotFoundHandler = handlers.ErrorRoute()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	fmt.Printf("🚀 Lancement de l'api sur le port %s\n", port)
	handler := c.Handler(router)

	err := http.ListenAndServe("0.0.0.0:"+port, handler)

	if err != nil {
		fmt.Printf("Err start serveur : %v\n", err)
	}
}
