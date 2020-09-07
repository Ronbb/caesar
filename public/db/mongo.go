package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	mops "go.mongodb.org/mongo-driver/mongo/options"
)

// Options 数据库选项
type Options struct {
	URL string
}

// New Create a mongodb client.
func New(options Options) (*mongo.Client, error) {
	// Set client options
	clientOptions := mops.Client().ApplyURI(options.URL)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil
}
