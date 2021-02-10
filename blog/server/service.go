package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/zemags/gRPSstudy/blog/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	driver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implement empty service with backward compatibility
type Service struct {
	pb.UnimplementedBlogServiceServer
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

// NewService make new empty Service
func NewService() *Service {
	return &Service{}
}

func createMongoClient(mongoURL string) (*driver.Client, error) {
	log.Printf("Create mongo client for %s", mongoURL)
	if mongoURL == "" {
		log.Fatalf("Cannot create client for Mongo db, no URL")
		return nil, errors.New("no url for Mongo client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := driver.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatalf("Cannot connect to %v", err)
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

// GetCollection impliment collection and return error
func (m *Mongo) getMongoCollection() *driver.Collection {
	log.Printf("Create mongo collection %v for db %v", m.Collection, m.DBName)
	coll := m.Database(m.DBName).Collection(m.Collection)
	return coll
}

// CreateBlog insert value for blog to db
func (m *Mongo) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
	blog := req.GetBlog()

	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	coll := m.Database(m.DBName).Collection(m.Collection)
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
