package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/zemags/gRPSstudy/calculator/calculator/pb"
)

func clientStreaming(c pb.CalculatorClient) {
	fmt.Println("Run client streaming with server")
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average %v", err)
	}
	for i := 1; i < 5; i++ {
		request := &pb.StreamRequest{
			PositiveInteger: int32(i),
		}
		stream.Send(request)
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receivinf response %v", err)
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
		log.Fatalf("error while callig rpc Decomposition %v", err)
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
