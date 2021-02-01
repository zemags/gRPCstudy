package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/zemags/gRPSstudy/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet function was invoked", req)
	firstName := req.GetGreeting().GetFirstName()
	result := fmt.Sprintf("Hello %s", firstName)
	response := &greetpb.GreetResponse{
		Result: result,
	}
	return response, nil

}

func main() {
	// write listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051") // 50051 default port for grpc
	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed to serve", err)
	}

}
