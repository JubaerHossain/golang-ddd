package main

import (
	"log"

	"github.com/JubaerHossain/golang-ddd/internal/core/server" // Import the generated gRPC client package
)

const (
	address = "localhost:50051" // Address of the gRPC server
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
