package main

import (
	"log"
	"net/http"
)

func main() {

	// 1. Initialize your store
	store := NewStore()

	// 2. Define your routes (POST /modifications, GET /modifications, GET /stats)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /modifications", store.handlePostMod)
	mux.HandleFunc("GET /modifications", store.handleGetMods)
	mux.HandleFunc("GET /stats", store.handleGetStats)

	// 3. Start the server on an open port of your choice
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
