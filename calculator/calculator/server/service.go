package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"

	"github.com/zemags/gRPSstudy/calculator/calculator/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implement empty service with backward compatibility
type Service struct {
	pb.UnimplementedCalculatorServer
}

// NewService make new empty Service
func NewService() *Service {
	return &Service{}
}

// SquareRoot implement grpc error
func (*Service) SquareRoot(ctx context.Context, req *pb.SquareRootRequest) (*pb.SquareRootResponse, error) {
	fmt.Println("Received SquareRoot RPC")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			// make own error. return nil as result and error
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number %v", number),
		)
	}
	// if all good return result and nil error
	return &pb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

// Maximum get stream of integers and send back maximum
func (*Service) Maximum(stream pb.Calculator_MaximumServer) error {
	fmt.Println("Run find maximum streaming")
	var currentValue int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
			return err
		}

		if req.PositiveInteger > currentValue {
			currentValue = req.PositiveInteger
			msg := &pb.ResponseMaximum{
				Maximum: currentValue,
			}
			if err := stream.Send(msg); err != nil {
				log.Fatalf("Error while sending to client stream %v", err)
				return err
			}
		}
	}
}

// Average get integers from client and send back average
func (*Service) Average(stream pb.Calculator_AverageServer) error {
	fmt.Println("Run average streaming")
	var sum int32
	var count float64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(
				&pb.ResponseDouble{
					Average: float64(sum) / count,
				})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}
		posInt := req.PositiveInteger
		sum += posInt
		count++
	}
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
