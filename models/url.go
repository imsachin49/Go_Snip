package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// URL struct represents the data stored in MongoDB
type URL struct {
	ID          string `bson:"_id,omitempty"`
	OriginalURL string `bson:"original_url"`
	ShortURL    string `bson:"short_url"`
}

// CreateURL inserts a new URL record into MongoDB
func CreateURL(collection *mongo.Collection, originalURL, shortURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url := URL{OriginalURL: originalURL, ShortURL: shortURL}
	_, err := collection.InsertOne(ctx, url)
	return err
}

// GetOriginalURL retrieves the original URL from MongoDB by the short URL code
func GetOriginalURL(collection *mongo.Collection, shortURL string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var url URL
	filter := bson.M{"short_url": shortURL}
	err := collection.FindOne(ctx, filter).Decode(&url)
	return url.OriginalURL, err
}
