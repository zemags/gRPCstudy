package main

import (
	"context"
	"log"

	"github.com/zemags/gRPSstudy/calculator/calculator/pb"
)

// Service implement empty service with backward compatibility
type Service struct {
	pb.UnimplementedCalculatorServer
}

// NewService make new empty Service
func NewService() *Service {
	return &Service{}
}

// Decomposition get one positive integer from client and make decomposition
// with printing factor values
func (*Service) Decomposition(req *pb.StreamRequest, stream pb.Calculator_DecompositionServer) error {
	log.Printf("Run new decomposition streaming")

	var k int32 = 2

	posInt := req.PositiveInteger
	for posInt > 1 {
		if posInt%k == 0 {
			res := &pb.StreamResponse{
				Factor: k,
			}
			stream.Send(res)
			posInt /= k
		} else {
			k++
		}
	}
	return nil
}

// Calc get two values from client and return sum of it
func (*Service) Calc(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	values := req.GetValues()
	firstValue := values.FirstValue
	secondValue := values.SecondValue

	response := &pb.CalculatorResponse{
		Sum: firstValue + secondValue,
	}
	return response, nil
}
