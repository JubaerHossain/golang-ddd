package main

import (
	"github.com/JubaerHossain/golang-ddd/internal/core/database/seed"
)

func main() {

	// Start the server
	if err := seed.NewSeed(); err != nil {
		// Handle error
		panic(err)
	}
}
