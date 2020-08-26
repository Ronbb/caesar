package db

import (
	"github.com/ronbb/caesar/public/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
  // database Mongodb database for caesar.
  database *mongo.Database
  // usersCol Mongodb collection with users' info.
  usersCol *mongo.Collection
)

func init()  {
  client, err := db.New(db.Options{
    URL: "mongodb://127.0.0.1:27017",
  })
  if err != nil {
    panic(err)
  }
  database = client.Database("caesar")
  usersCol = database.Collection("users")
}
