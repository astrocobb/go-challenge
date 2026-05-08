package main

import (
	"errors"
)

type Modification struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Room     string `json:"room"`
	Cost     int    `json:"cost"`
	Priority int    `json:"priority"`
}

// validate validates the modification
func (mod *Modification) validate() error {
	if mod.ID <= 0 {
		return errors.New("id must be greater than 0")
	}
	if mod.Name == "" {
		return errors.New("name is required")
	}
	if mod.Room == "" {
		return errors.New("room is required")
	}
	if mod.Cost < 0 {
		return errors.New("cost must be greater than or equal to 0")
	}
	if mod.Priority < 1 || mod.Priority > 5 {
		return errors.New("priority must be between 1 and 5")
	}
	return nil
}
