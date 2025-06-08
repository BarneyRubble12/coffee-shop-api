package model

// Coffee represents a coffee product in the shop
type Coffee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Origin string  `json:"origin"`
	Price  float64 `json:"price"`
}
