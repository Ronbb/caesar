package db

import (
	"github.com/ronbb/caesar/public/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
  // Database Mongodb database for caesar.
	Database *mongo.Database
)

func init()  {
  client, err := db.New(db.Options{
    URL: "mongodb://127.0.0.1:27017",
  })
  if err != nil {
    panic(err)
  }
  Database = client.Database("caesar")
}
