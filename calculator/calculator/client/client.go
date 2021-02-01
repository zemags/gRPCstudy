package main

import (
	"context"
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
	request := getRequest(client)

	response, err := client.Calc(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while Cacl RPC %v", err)
	}
	log.Fatalf("Answer responce %v", response.Sum)

}

func getRequest(client pb.CalculatorClient) *pb.CalculatorRequest {
	request := &pb.CalculatorRequest{
		Values: &pb.Values{
			FirstValue:  1,
			SecondValue: 2,
		},
	}

	return request

}
