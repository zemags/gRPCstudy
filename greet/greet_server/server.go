package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// write listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051") // 50051 default port for grpc
	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	s := grpc.NewServer()
}
