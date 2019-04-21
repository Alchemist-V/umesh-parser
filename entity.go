package main

import "fmt"

// Entity represents DB entity for DSR record.
type Entity struct {
	ID          string
	Description string
	Unit        string
	Rate        float64
}

// Reciver to print an entity.
func (e Entity) print() {
	fmt.Printf("| %s | %s | %s | %f |\n", e.ID, e.Description, e.Unit, e.Rate)
}
