package main

import (
	"log"
	"net/http"

	"receipt-processor/internal/api"
)

func main() {
	// Create a new router
	router := api.NewRouter()

	// Start the server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
