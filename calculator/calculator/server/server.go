package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/zemags/gRPSstudy/calculator/calculator/pb"

	"google.golang.org/grpc"
)

type Service struct {
	pb.UnimplementedCalculatorServer
}

func (*Service) Calc(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	values := req.GetValues()
	firstValue := values.FirstValue
	secondValue := values.SecondValue

	response := &pb.CalculatorResponse{
		Sum: firstValue + secondValue,
	}
	return response, nil
}

func main() {
	fmt.Println("Calculator ready.")

	// create listener
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterCalculatorServer(server, new(Service))
	log.Printf("Running server on %s", lis.Addr().String())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server listener %v", err)
	}

}
