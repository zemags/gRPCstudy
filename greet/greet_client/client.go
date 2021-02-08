package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/zemags/gRPSstudy/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Client executed")

	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // withinsecure = without ssl

	certFile := "ssl/ca.crt" // Certificate Authority trust certificate
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	// with ssl
	if sslErr != nil {
		log.Fatalln("Error while loading CA trust certificate ", sslErr)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln("Could not connect", err)
	}
	defer conn.Close()

	// create client
	c := greetpb.NewGreetServiceClient(conn)
	doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBiDirectionalStreaming(c)
	// doUnaryDeadline(c, 1*time.Second) // should complete
	// doUnaryDeadline(c, 5*time.Second) // should timeout

}

func doUnaryDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("Run deadline unary rpc")

	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "john",
			LastName:  "dor",
		},
	}

	// sleep 5 sec until deadline
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	response, err := c.GreetDeadline(ctx, request)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// if deadline exceeded
			if respErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit. Deadline was exceeded.")
			} else {
				fmt.Println("unexpected error ", respErr)
			}
		} else {
			log.Fatalln("Error while GreetDeadline RPC", err)
		}
		return
	}
	log.Fatalln("Response from GreetDeadline", response.Result)

}

func doBiDirectionalStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Client bidi stream")

	stream, err := c.GreetEvery(context.Background())
	if err != nil {
		log.Fatalln("Error while calling GreetEvery", err)
		return
	}

	waitChan := make(chan struct{})
	// send bunch of msg async
	go func() {
		for i := 0; i < 5; i++ {
			request := &greetpb.GreetEveryRequest{
				Greeting: &greetpb.Greeting{
					FirstName: fmt.Sprintf("John %d", i),
				},
			}
			stream.Send(request)
		}
		stream.CloseSend()
	}()

	// receive a bunch of msg
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				// end of the stream close channel
				break
			}
			if err != nil {
				log.Fatalln("Error while receiving", err)
			}
			fmt.Println("Received ", res.GetResult())
		}
		close(waitChan)
	}()

	// block until everything is done
	<-waitChan
}

func doClientStreaming(c greetpb.GreetServiceClient) {
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
		log.Fatalln("Error while Greet RPC", err)
	}
	log.Fatalln("Response from Greet", response.Result)
}
