package main

import (
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
	fmt.Println("Created client", c)
}
