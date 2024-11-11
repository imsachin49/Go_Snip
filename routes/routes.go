package routes

import (
	"github.com/gorilla/mux"
	"github.com/imsachin49/URL_Shortener/handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(router *mux.Router, collection *mongo.Collection) {
	router.HandleFunc("/shorten", handlers.ShortenURLHandler(collection)).Methods("POST")
	router.HandleFunc("/{shortURL}", handlers.RedirectURLHandler(collection)).Methods("GET")
}
