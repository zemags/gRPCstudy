package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zemags/gRPSstudy/blog/pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client BlogService executed")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	blog := &pb.Blog{
		AuthorId: "John",
		Title:    "How to build gRPS servece",
		Content:  "First of all, you need...",
	}

	res, err := c.CreateBlog(context.Background(), &pb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Cannot create blog %v", err)
	}
	fmt.Printf("Blog was created %v", res)

}
