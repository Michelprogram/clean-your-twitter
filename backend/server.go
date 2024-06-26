package main

import (
	"api-clean-twitter/cron"
	"api-clean-twitter/database"
	h "api-clean-twitter/handlers"
	m "api-clean-twitter/middleware"
	"flag"
	"log"
	"net/http"

	"fmt"

	"github.com/rs/cors"
	"golang.org/x/net/websocket"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {

	var err error

	var port string

	flag.StringVar(&port, "port", "3021", "Port to listen on")

	_ = godotenv.Load(".env")

	err = database.OpenConnection()

	if err != nil {
		log.Fatalf("Connection to the database : %s", err)
	}

	router := mux.NewRouter()

	router.Use(m.Logs)

	//Authentification with twitter
	twitter_routes := router.PathPrefix("/twitter").Subrouter()
	twitter_routes.HandleFunc("/auth", h.AuthentificationTwitter).Methods("GET")

	//Interaction with backend
	backend_routes := router.PathPrefix("/backend").Subrouter()
	backend_routes.Use(m.AuthBackend)

	backend_routes.HandleFunc("/auth", h.AuthentificationBackend).Methods("GET")
	backend_routes.HandleFunc("/clean", h.CleanTweets).Methods("POST")

	websocket_routes := router.PathPrefix("/ws").Subrouter()

	websocket_routes.Handle("/tweets", websocket.Handler(h.Fetch))

	//Ping server
	test_routes := router.PathPrefix("/test").Subrouter()
	test_routes.Use(m.AuthBackend)

	test_routes.HandleFunc("/", h.TestCode).Methods("GET")
	test_routes.HandleFunc("/ping", h.Test).Methods("GET")

	//router.NotFoundHandler = handlers.ErrorRoute() update
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://clean-your-tw.online"},
		AllowCredentials: true,
	})

	go cron.CleanTweets()

	fmt.Printf("🚀 Lancement de l'api sur le port %s\n", port)
	handler := c.Handler(router)
	c.Handler(router)

	err = http.ListenAndServe("0.0.0.0:"+port, handler)

	if err != nil {
		fmt.Printf("Err start server : %v\n", err)
	}
}
