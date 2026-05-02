package main

import "sync"

type Store struct {
	mu            sync.Mutex
	modifications []Modification
}

type Stats struct {
	TotalMods int `json:"totalModifications"`
	TotalCost int `json:"totalCost"`
}
