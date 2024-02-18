package main

import (
	"github.com/JubaerHossain/golang-ddd/internal/core/server"
)

func main() {
	
	// Create a new server instance
	srv := server.NewServer()
	// Start the server
	if err := srv.Start(); err != nil {
		// Handle error
		panic(err)
	}
}
