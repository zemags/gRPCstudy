package main

import (
	"context"
	"fmt"

	"github.com/zemags/gRPSstudy/blog/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	positive "go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implement empty service with backward compatibility and collect all opts
type Service struct {
	pb.UnimplementedBlogServiceServer
	Mongo *Mongo
}

type blogItem struct {
	ID       positive.ObjectID `bson:"_id,omitempty"`
	AuthorID string            `bson:"author_id"`
	Title    string            `bson:"title"`
	Content  string            `bson:"content"`
}

// NewService make new Service
func NewService(m *Mongo) *Service {
	return &Service{
		Mongo: m,
	}
}

// CreateBlog insert value for blog to db
func (s *Service) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
	fmt.Println("CreateBlog rpc service")
	blog := req.GetBlog()

	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	coll := s.Mongo.Database(s.Mongo.DBName).Collection(s.Mongo.Collection)
	res, err := coll.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, // Internal db error
			fmt.Sprintf("Internal Error %v", err),
		)
	}
	objId, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal, // Internal db error
			fmt.Sprintf("Can't convert to objectID"),
		)
	}

	return &pb.CreateBlogResponse{
		Blog: &pb.Blog{
			Id:       objId.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil

}

// ReadBlog return blog id exist else NOT_FOUND
func (s *Service) ReadBlog(ctx context.Context, req *pb.ReadBlogRequest) (*pb.ReadBlogResponse, error) {
	fmt.Println("ReadBlog rpc service")

	blogId := req.GetBlogId()
	objId, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse Id"),
		)
	}

	data := &blogItem{}
	coll := s.Mongo.Database(s.Mongo.DBName).Collection(s.Mongo.Collection)
	filter := bson.M{"_id": objId}
	res := coll.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound, // grpc not found error
			fmt.Sprintf("Cannot find blog with specified ID %v", err),
		)
	} // pass value from res to data
	return &pb.ReadBlogResponse{
		Blog: &pb.Blog{
			Id:       data.AuthorID,
			AuthorId: data.AuthorID,
			Content:  data.Content,
			Title:    data.Title,
		},
	}, nil
}

// UpdateBlog update blog poset by getting blog id
func (s *Service) UpdateBlog(ctx context.Context, req *pb.UpdateBlogRequest) (*pb.UpdateBlogResponse, error) {
	fmt.Println("UpdateBlog rpc server")

	blog := req.GetBlog()
	objId, err := primitive.ObjectIDFromHex(blog.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse Id"),
		)
	}

	data := &blogItem{}
	filter := bson.M{"_id": objId}
	coll := s.Mongo.Database(s.Mongo.DBName).Collection(s.Mongo.Collection)
	res := coll.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound, // grpc not found error
			fmt.Sprintf("Cannot find blog with specified ID %v", err),
		)
	} // pass value from res to data

	data.AuthorID = blog.GetAuthorId()
	data.Title = blog.GetTitle()
	data.Content = blog.GetContent()

	_, errUpdt := coll.ReplaceOne(context.Background(), filter, data)
	if errUpdt != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot replace blog with specified ID %v", errUpdt),
		)
	}
	return &pb.UpdateBlogResponse{
		Blog: &pb.Blog{
			Id:       data.AuthorID,
			AuthorId: data.AuthorID,
			Content:  data.Content,
			Title:    data.Title,
		},
	}, nil

}

// DeleteBlog remove blog post from Mongo by
func (s *Service) DeleteBlog(ctx context.Context, req *pb.DeleteBlogRequest) (*pb.DeleteBlogResponse, error) {
	fmt.Println("DeleteBlog rpc service")
	objId, err := primitive.ObjectIDFromHex(req.GetBlogId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse Id"),
		)
	}

	filter := bson.M{"_id": objId}
	coll := s.Mongo.Database(s.Mongo.DBName).Collection(s.Mongo.Collection)
	_, errDel := coll.DeleteOne(context.Background(), filter)
	if errDel != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete blog with specified ID %v", errDel),
		)
	}
	return &pb.DeleteBlogResponse{BlogId: req.GetBlogId()}, nil
}
