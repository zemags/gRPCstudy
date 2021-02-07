package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/zemags/gRPSstudy/greet/greetpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet function was invoked", req)
	firstName := req.GetGreeting().GetFirstName()
	result := fmt.Sprintf("Hello %s", firstName)
	response := &greetpb.GreetResponse{
		Result: result,
	}
	return response, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Println("Run server streaming")
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := fmt.Sprintf("Hello %s %d", firstName, i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("Run client streaming")
	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{ // return result or error
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("error while reading client stream %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += firstName + "\n"
	}
}

func (*server) GreetEvery(stream greetpb.GreetService_GreetEveryServer) error {
	fmt.Println("Run client streaming and replay")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalln("Error while reading client stream", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "hello " + firstName
		err = stream.Send(&greetpb.GreetEveryResponse{
			Result: result,
		})
		if err != nil {
			log.Fatalln("error while sending data to client", err)
			return err
		}
	}
}

func (*server) GreetDeadline(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("GreetDeadline function was invoked", req)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			// the client canceled the request
			fmt.Println("The client cancel the request")
			return nil, status.Error(codes.Canceled, "the client cacnceleed the request!")
		}
		time.Sleep(1 * time.Second)
	}

	firstName := req.GetGreeting().GetFirstName()
	result := fmt.Sprintf("Hello %s", firstName)
	response := &greetpb.GreetResponse{
		Result: result,
	}
	return response, nil
}

func main() {
	// write listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051") // 50051 default port for grpc
	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed to serve", err)
	}

}
