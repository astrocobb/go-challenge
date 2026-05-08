package main

import (
	"errors"
	"sync"
)

type Store struct {
	mu            sync.Mutex
	modifications []Modification
}

type Stats struct {
	TotalMods int `json:"totalModifications"`
	TotalCost int `json:"totalCost"`
}

// newStore Constructor
func newStore() *Store {
	return &Store{
		modifications: make([]Modification, 0),
	}
}

// addMods adds a modification to the store calling it
func (store *Store) addMods(mod Modification) error {
	store.mu.Lock()
	defer store.mu.Unlock()
	for _, existingMod := range store.modifications {
		if existingMod.ID == mod.ID {
			return errors.New("modification already exists")
		}
	}
	store.modifications = append(store.modifications, mod)
	return nil
}

// getMods returns all the store's modifications
func (store *Store) getMods() []Modification {
	store.mu.Lock()
	defer store.mu.Unlock()
	out := make([]Modification, len(store.modifications))
	copy(out, store.modifications)
	return out
}

// getStats returns the store's stats
func (store *Store) getStats() Stats {
	store.mu.Lock()
	defer store.mu.Unlock()
	var totalCost int
	for _, mod := range store.modifications {
		totalCost += mod.Cost
	}
	return Stats{TotalMods: len(store.modifications), TotalCost: totalCost}
}
