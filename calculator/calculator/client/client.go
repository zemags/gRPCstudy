package main

import (
	"log"

	"github.com/zemags/gRPSstudy/calculator/calculator/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)
	// getRequest(client)
	// serverStreaming(client)
	// clientStreaming(client)
	// biDirectionalStreaming(client)
	errorUnary(client)
}
