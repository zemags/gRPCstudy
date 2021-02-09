package main

import (
	"context"
	"fmt"

	"github.com/zemags/gRPSstudy/blog/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implement empty service with backward compatibility
type Service struct {
	pb.UnimplementedBlogServiceServer
}

// NewService make new empty Service
func NewService() *Service {
	return &Service{}
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` // autoincrement
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

// CreateBlog insert value for blog to db
func (*Service) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
	blog := req.GetBlog()

	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, // Internal db error
			fmt.Sprintf("Internal Error %v", err),
		)
	}
	objID, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal, // Internal db error
			fmt.Sprintf("Can't convert to objectID"),
		)
	}

	return &pb.CreateBlogResponse{
		Blog: &pb.Blog{
			Id:       objID.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil

}
