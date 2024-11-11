package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"url-shortener/routes"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // MongoDB connection setup
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)

    // Get the URL collection
    collection := client.Database("urlshortener").Collection("urls")

    // Set up the router
    router := mux.NewRouter()
    routes.SetupRoutes(router, collection)

    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
