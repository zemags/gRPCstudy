package main

import (
	"context"
	"fmt"
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
	doUnary(c)

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
		log.Fatalf("Error while Greep RPC", err)
	}
	log.Fatalf("Response from Greet", response.Result)
}
