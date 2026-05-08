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

// NewStore Constructor
func NewStore() *Store {
	return &Store{
		modifications: make([]Modification, 0),
	}
}

// AddMods adds a modification to the store calling it
func (store *Store) AddMods(mod Modification) error {
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

// GetMods returns all the store's modifications
func (store *Store) GetMods() []Modification {
	store.mu.Lock()
	defer store.mu.Unlock()
	out := make([]Modification, len(store.modifications))
	copy(out, store.modifications)
	return out
}

// GetStats returns the store's stats
func (store *Store) GetStats() Stats {
	store.mu.Lock()
	defer store.mu.Unlock()
	var totalCost int
	for _, mod := range store.modifications {
		totalCost += mod.Cost
	}
	return Stats{TotalMods: len(store.modifications), TotalCost: totalCost}
}
