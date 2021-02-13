package main

import (
	"context"
	"errors"
	"log"
	"time"

	driver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

// CreateMongoClient connect to Mongo DB and return client
func CreateMongoClient(mongoURL string) (*driver.Client, error) {
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

// CreateMongoServer create Mongo struct with implemented client and options
func CreateMongoServer(client *driver.Client, mp MongoParams) *Mongo {
	log.Printf("Create new mongo server")
	return &Mongo{
		Client:      client,
		MongoParams: mp,
	}
}
