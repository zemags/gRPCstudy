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
		Title:    "How to build gRPC servece",
		Content:  "First of all, you need...",
	}

	createReq, errCreate := c.CreateBlog(context.Background(), &pb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Cannot create blog %v", errCreate)
	}
	fmt.Printf("Blog was created %v", createReq)

	blogId := createReq.GetBlog().GetId()

	// read blog
	fmt.Println("Read blog client")
	_, errRead := c.ReadBlog(context.Background(), &pb.ReadBlogRequest{BlogId: "1232aa"}) // error check
	if errRead != nil {
		fmt.Printf("Cannot read blog: %v \n", errRead)
	}

	readReq, errRead := c.ReadBlog(context.Background(), &pb.ReadBlogRequest{BlogId: blogId})
	if errRead != nil {
		fmt.Printf("Cannot read blog: %v \n", errRead)
	}
	fmt.Printf("Blog was read %v", readReq)

	// update blog
	fmt.Println("Update blog client")
	updtBlog := &pb.Blog{
		Id:       blogId,
		AuthorId: "John",
		Title:    "How to build gRPC servece Part 2",
		Content:  "Second of all, you need...",
	}

	updtReq, errUpdt := c.UpdateBlog(context.Background(), &pb.UpdateBlogRequest{Blog: updtBlog})
	if errUpdt != nil {
		fmt.Printf("Cannot update blog: %v \n", errUpdt)
	}
	fmt.Printf("Blog was update %v", updtReq)

	// delete blog
	fmt.Println("Delete blog client")
	_, errDel := c.DeleteBlog(context.Background(), &pb.DeleteBlogRequest{BlogId: "1232aa"}) // error check
	if errDel != nil {
		fmt.Printf("Cannot delete blog: %v \n", errDel)
	}

	deleteReq, errDel := c.DeleteBlog(context.Background(), &pb.DeleteBlogRequest{BlogId: blogId})
	if errDel != nil {
		fmt.Printf("Cannot delete blog: %v \n", errDel)
	}
	fmt.Printf("Blog was delete %v", deleteReq)
}
