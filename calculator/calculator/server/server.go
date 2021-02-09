package main

import (
	"fmt"
	"log"
	"net"

	"github.com/zemags/gRPSstudy/calculator/calculator/pb"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator ready.")

	// create listener
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterCalculatorServer(server, NewService())

	// Register reflection service on gRPC server.
	reflection.Register(server)

	log.Printf("Running server on %s", lis.Addr().String())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server listener %v", err)
	}
}
