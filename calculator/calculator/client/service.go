package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/zemags/gRPSstudy/calculator/calculator/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func errorUnary(c pb.CalculatorClient) {
	fmt.Println("Run SquareRoot rpc")

	res, err := c.SquareRoot(context.Background(), &pb.SquareRootRequest{Number: int32(-2)})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual err from gRPC (user error)
			fmt.Printf("Error message from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code()) // InvalidArgument
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("Negative number send")
				return
			}
		} else {
			log.Fatalf("Bog error calling SquareRoot %v", err)
			return
		}
	}
	fmt.Printf("Result %v", res.GetNumberRoot())
}

func biDirectionalStreaming(c pb.CalculatorClient) {
	fmt.Println("Run bidi client streaming with server")

	stream, err := c.Maximum(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Maximum rpc %v", err)
		return
	}

	sliceOfInts := []int32{1, 5, 3, 6, 2, 20}
	// send messages to server
	for _, i := range sliceOfInts {
		request := &pb.StreamRequest{
			PositiveInteger: i,
		}
		stream.Send(request)
	}
	stream.CloseSend()

	// receive messages from server
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving %v", err)
		}
		fmt.Println("Received ", res.GetMaximum())
	}
}

func clientStreaming(c pb.CalculatorClient) {
	fmt.Println("Run client streaming with server")
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average rpc %v", err)
	}
	for i := 1; i < 5; i++ {
		request := &pb.StreamRequest{
			PositiveInteger: int32(i),
		}
		stream.Send(request)
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response %v", err)
	}
	fmt.Printf("Average: %v\n", response)
}

func serverStreaming(c pb.CalculatorClient) {
	fmt.Println("Run server streaming with client")
	request := &pb.StreamRequest{
		PositiveInteger: 120,
	}
	streamResult, err := c.Decomposition(context.Background(), request)
	if err != nil {
		log.Fatalf("error while callig Decomposition rpc %v", err)
	}
	for {
		msg, err := streamResult.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream %v", err)
		}
		fmt.Printf("Current factor number is: %d\n", msg.GetFactor())
	}
}

func getRequest(client pb.CalculatorClient) {
	request := &pb.CalculatorRequest{
		Values: &pb.Values{
			FirstValue:  1,
			SecondValue: 2,
		},
	}
	response, err := client.Calc(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while Cacl RPC %v", err)
	}
	log.Fatalf("Answer responce %v", response.Sum)
}
