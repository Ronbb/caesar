package db

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func db() {
	fmt.Println("hello world")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
  defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
  }

  runTest(client)
}

func runTest(client *mongo.Client)  {
  ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
  defer cancel()
  c, err := client.Database("iot").Collection("iot").Find(ctx, bson.D{})
  // ss, err := client.Database("iot").ListCollectionNames(ctx, bson.D{})
  if err != nil {
		fmt.Println(err.Error())
		panic(err)
  }
  if c.Err() != nil {
		fmt.Println(c.Err())
		panic(c.Err())
  }

  var m map[string]interface{}

  for c.Next(ctx) {
    c.Decode(&m)
    b, _ :=json.Marshal(m)
    fmt.Println(string(b))
  }
}
