package db

import (
	"context"
	"errors"
	"log"

	"github.com/ronbb/caesar/public/models/caesar"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	username = "username"
)

// UserSignUp Add a user to db.
func UserSignUp(user *caesar.UserRegistration) error {
	r := usersCol.FindOne(context.TODO(), bson.M{
		username: user.Username,
	})

	userInfo := &caesar.UserInfo{}
	err := r.Decode(userInfo)

	if err == nil {
    err = errors.New("duplicate username")
		log.Println(err.Error())
		return err
  }

	if err.Error() != mongo.ErrNoDocuments.Error() {
		log.Println(err.Error())
		return err
	}

	_, err = usersCol.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// UserSignIn Check a user from db.
func UserSignIn(user *caesar.Authentication) (*caesar.UserInfo, error) {
	res := usersCol.FindOne(context.TODO(), user)
	userInfo := &caesar.UserInfo{}
	err := res.Decode(userInfo)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return userInfo, nil
}
