package main

import (
	"fmt"
	"net/http"
)

func main() {

	// 1. Initialize your store
	store := Store{}

	// 2. Define your routes (POST /modifications, GET /modifications)
	http.HandleFunc("POST /modifications", store.postModification)
	http.HandleFunc("GET /modifications", store.getModifications)
	http.HandleFunc("GET /modifications/stats", store.getStoreStats)

	// 3. Start the server on an open port of your choice
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
