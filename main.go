package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tiagosantini/shortgo/internal/models"
	"github.com/tiagosantini/shortgo/internal/shortener"
)

func shortenUrlHandler(w http.ResponseWriter, req *http.Request) {
	// Process request body
	var originalUrl models.OriginalUrl;

	err := json.NewDecoder(req.Body).Decode(&originalUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req.Body.Close()

	// Generate shortened URL
	shortenedUrl, err := shortener.ShortenUrl(&originalUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode URL data as JSON and write to response stream
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(shortenedUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func executeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		// Set application/json as default Content-Type
	    w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle(
		"POST /shorten",
		executeMiddleware(http.HandlerFunc(shortenUrlHandler)),
	)

	log.Print("Server listening on port 8080.")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
