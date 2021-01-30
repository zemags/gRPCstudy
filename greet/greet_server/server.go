package main

import (
	"log"
	"net"

	"github.com/zemags/gRPSstudy/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	// write listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051") // 50051 default port for grpc
	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

}
