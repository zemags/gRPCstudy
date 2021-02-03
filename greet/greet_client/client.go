package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/zemags/gRPSstudy/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client executed")

	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // withinsecure = without ssl
	if err != nil {
		log.Fatalln("Could not connect", err)
	}
	defer con.Close()

	// create client
	c := greetpb.NewGreetServiceClient(con)
	// doUnary(c)

	// doServerStreaming(c)
	doClientStreamin(c)

}

func doClientStreamin(c greetpb.GreetServiceClient) {
	request := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Dor",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Bor",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Tor",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalln("Error while calling LongGreet", err)
	}

	// send message individually
	for _, req := range request {
		fmt.Println(req)
		stream.Send(req)
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error while receiving response", err)
	}
	fmt.Println("Long greet response", response)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Run server streaming with client")
	request := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "John",
			LastName:  "dor",
		},
	}
	streamResult, err := c.GreetManyTimes(context.Background(), request)
	if err != nil {
		log.Fatalln("Error while calling rpc GreetManyTimes", err)
	}
	for {
		msg, err := streamResult.Recv()
		if err == io.EOF {
			// reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalln("Error while reading stream", err)
		}
		fmt.Println("Response from GreetManyTimes", msg.GetResult())
	}
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Run unary RPC")
	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "john",
			LastName:  "dor",
		},
	}
	response, err := c.Greet(context.Background(), request)
	if err != nil {
		log.Fatalln("Error while Greep RPC", err)
	}
	log.Fatalln("Response from Greet", response.Result)
}
