package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// handlePostMod handles the POST /modifications route
func (store *Store) handlePostMod(w http.ResponseWriter, r *http.Request) {
	var req Modification
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := req.validate(); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := store.addMods(req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, req)
}

// handleGetMods handles the GET /modifications route
func (store *Store) handleGetMods(w http.ResponseWriter, r *http.Request) {
	mods := store.getMods()
	writeJSON(w, http.StatusOK, mods)
}

// handleGetStats handles the GET /stats route
func (store *Store) handleGetStats(w http.ResponseWriter, r *http.Request) {
	stats := store.getStats()
	writeJSON(w, http.StatusOK, stats)
}

// writeJSON helper function to write JSON to the response
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("json encode error: %v", err)
	}
}
