package main

type Modification struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Room     string `json:"room"`
	Cost     int    `json:"cost"`
	Priority int    `json:"priority"`
}
