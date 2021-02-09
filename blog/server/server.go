package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/zemags/gRPSstudy/blog/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection

func main() {
	// if code crashed get filename and line number of code
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Starting BlogService")

	// connect to mongo db
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	// if db doesnt exist will be created, Collection - like table in psql
	collection = client.Database("mydb").Collection("blog")

	lsn, lsnErr := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", lsnErr)
	}
	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	pb.RegisterBlogServiceServer(server, NewService())

	log.Printf("Running server on %s", lsn.Addr().String())

	go func() {
		fmt.Println("Starting server")
		if err := server.Serve(lsn); err != nil {
			log.Fatalf("Failed to serve listener %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	server.Stop()
	lsn.Close()
	fmt.Println("Done")
}
