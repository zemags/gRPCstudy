package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/zemags/gRPSstudy/blog/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	driver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implement empty service with backward compatibility
type Service struct {
	pb.UnimplementedBlogServiceServer
	Mongo *Mongo
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` // autoincrement
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

// Mongo collect all options to work with mongo
type Mongo struct {
	*driver.Client
	MongoParams
}

// MongoParams collection given dbname and collection
type MongoParams struct {
	DBName     string
	Collection string
}

// NewService make new Service
func NewService(m *Mongo) *Service {
	return &Service{
		Mongo: m,
	}
}

func createMongoClient(mongoURL string) (*driver.Client, error) {
	log.Printf("Create mongo client for %s", mongoURL)
	if mongoURL == "" {
		log.Fatalf("Cannot create client for Mongo db, no URL")
		return nil, errors.New("no url for Mongo client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := driver.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatalf("Cannot initialize connection %v", err)
		return nil, errors.New("cannot connect to MongoDB")
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Cannot connect %v", err)
		return nil, errors.New("cannot connect to MongoDB")
	}

	return client, nil
}

func createMongoServer(client *driver.Client, mp MongoParams) *Mongo {
	log.Printf("Create new mongo server")
	return &Mongo{
		Client:      client,
		MongoParams: mp,
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

// ReadBlog return blog id exist else NOT_FOUND
func (s *Service) ReadBlog(ctx context.Context, req *pb.ReadBlogRequest) (*pb.ReadBlogResponse, error) {
	fmt.Println("ReadBlog rpc service")

	blogID := req.GetBlogId()
	objID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		log.Fatalf("Cannot create objectId from hex string %v", err)
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	data := &blogItem{}
	coll := s.Mongo.Database(s.Mongo.DBName).Collection(s.Mongo.Collection)
	filter := bson.M{"_id": objID}
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
