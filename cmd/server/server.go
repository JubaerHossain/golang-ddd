package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
    pb "github.com/JubaerHossain/golang-ddd/internal/infrastructure/grpc" // Import the generated gRPC client package

)

const (
	address = "localhost:50051" // Address of the gRPC server
)

func main() {
	// Set up transport credentials for insecure connection
	creds := credentials.NewInsecure()

	// Set up a connection to the server with transport credentials
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewFoodAppClient(conn)

	// Example call to the gRPC server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.SomeRPCMethod(ctx, &pb.SomeRequest{ /* Provide request data */ })
	if err != nil {
		log.Fatalf("Error calling gRPC method: %v", err)
	}
	// Process the response from the server
	fmt.Println("Response:", response)
}
