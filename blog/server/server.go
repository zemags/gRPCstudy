package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/zemags/gRPSstudy/blog/pb"
	"google.golang.org/grpc"
)

// ServerOpts Define server options
// include server and db options
type ServerOpts struct {
	IPPort   string
	MongoURL string
}

// var coll *mongo.Collection

func main() {
	// if code crashed get filename and line number of code
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting BlogService server")

	servOpts := &ServerOpts{
		IPPort:   "localhost:50051",
		MongoURL: "mongodb://localhost:27017",
	}

	client, err := createMongoClient(servOpts.MongoURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	mgParams := MongoParams{DBName: "mydb", Collection: "blog"}
	mg := createMongoServer(client, mgParams)
	service := NewService(mg)

	lsn, lsnErr := net.Listen("tcp", servOpts.IPPort)
	if err != nil {
		log.Fatalf("Failed to listen %v", lsnErr)
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	pb.RegisterBlogServiceServer(server, service)

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
