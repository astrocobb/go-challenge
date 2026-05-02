package main

import (
	"encoding/json"
	"net/http"
)

func (store *Store) postModification(w http.ResponseWriter, r *http.Request) {

	store.mu.Lock()
	defer store.mu.Unlock()

	// decode the modification from the request body
	var mod Modification
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&mod)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// validate the modification doesn't exist
	for _, m := range store.modifications {
		if m.ID == mod.ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// validate the modification's priority
	if mod.Priority < 1 || mod.Priority > 5 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// add the modification to the store
	store.modifications = append(store.modifications, mod)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mod)
}

func (store *Store) getModifications(w http.ResponseWriter, r *http.Request) {

	store.mu.Lock()
	defer store.mu.Unlock()

	var status int
	if len(store.modifications) == 0 {
		status = http.StatusNoContent
	} else {
		status = http.StatusOK
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(store.modifications)
}

func (store *Store) getStoreStats(w http.ResponseWriter, r *http.Request) {

	store.mu.Lock()
	defer store.mu.Unlock()

	var count = len(store.modifications)
	var total int

	// loop through modifications, adding up the cost, and counting each
	for _, mod := range store.modifications {
		total += mod.Cost
	}

	// create a stats object
	stats := Stats{count, total}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stats)
}
