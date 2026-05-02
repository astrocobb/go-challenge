package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Modification struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Room     string `json:"room"`
	Cost     int    `json:"cost"`
	Priority int    `json:"priority"`
}

type Store struct {
	mu            sync.Mutex
	modifications []Modification
}

type Stats struct {
	TotalMods int `json:"totalModifications"`
	TotalCost int `json:"totalCost"`
}

func main() {

	// 1. Initialize your store
	store := Store{
		mu:            sync.Mutex{},
		modifications: []Modification{},
	}

	// 2. Define your routes (POST /modifications, GET /modifications)
	http.HandleFunc("POST /modifications", store.postModification)
	http.HandleFunc("GET /modifications", store.getModifications)
	http.HandleFunc("GET /modifications/stats", store.getStoreStats)

	// 3. Start the server on an open port of your choice
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func (store *Store) postModification(w http.ResponseWriter, r *http.Request) {

	// lock the store
	store.mu.Lock()
	defer store.mu.Unlock()

	// decode the modification from the request body
	var mod Modification
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&mod)

	// add the modification to the store
	store.modifications = append(store.modifications, mod)

	// write the response
	w.WriteHeader(http.StatusCreated)
	fmt.Println(mod, "successfully stored.")
}

func (store *Store) getModifications(w http.ResponseWriter, r *http.Request) {

	// lock the store
	store.mu.Lock()
	defer store.mu.Unlock()

	// write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(store.modifications)
}

func (store *Store) getStoreStats(w http.ResponseWriter, r *http.Request) {

	// lock the store
	store.mu.Lock()
	defer store.mu.Unlock()

	// declare variables
	var count int
	var total int

	// loop through modifications, adding up the cost, and counting each
	for _, mod := range store.modifications {
		count++
		total += mod.Cost
	}

	// create a stats object
	stats := Stats{count, total}

	// write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stats)
}
