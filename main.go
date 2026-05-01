package main

import (
	"encoding/json"
	"net/http"
	"strconv"
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

func main() {
	// 1. Initialize your store
	// store := new(Store)

	// 2. Define your routes (POST /modifications, GET /modifications)
	mux := http.NewServeMux()
	http.HandleFunc("POST /modifications", postModifications)
	mux.HandleFunc("GET /modifications", getModifications)

	// 3. Start the server on an open port of your choice
	http.ListenAndServe(":8080", mux)
}

func postModifications(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	body := r.Body

	if method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func getModifications(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	method := r.Method
	if method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO
	store := new(Store)
	modification := findModification(id, store)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(modification)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func findModification(id int, store *Store) Modification {
	var modification Modification
	for _, mod := range store.modifications {
		if mod.ID == id {
			modification = mod
		}
	}
	return modification
}
