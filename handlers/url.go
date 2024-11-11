package handlers

import (
    "encoding/json"
    "math/rand"
    "net/http"
    "url-shortener/models"

    "go.mongodb.org/mongo-driver/mongo"
)

// ShortenURLHandler creates a shortened URL
func ShortenURLHandler(collection *mongo.Collection) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var request struct{ URL string `json:"url"` }
        json.NewDecoder(r.Body).Decode(&request)

        shortURL := generateShortURL()
        err := models.CreateURL(collection, request.URL, shortURL)
        if err != nil {
            http.Error(w, "Could not create short URL", http.StatusInternalServerError)
            return
        }

        response := map[string]string{"short_url": shortURL}
        json.NewEncoder(w).Encode(response)
    }
}

// RedirectURLHandler redirects to the original URL
func RedirectURLHandler(collection *mongo.Collection) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        shortURL := vars["shortURL"]

        originalURL, err := models.GetOriginalURL(collection, shortURL)
        if err != nil {
            http.Error(w, "URL not found", http.StatusNotFound)
            return
        }

        http.Redirect(w, r, originalURL, http.StatusFound)
    }
}
