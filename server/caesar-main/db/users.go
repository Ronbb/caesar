package db

import (
	"context"
	"log"

	"github.com/ronbb/caesar/public/models/caesar"
)

// UserSignUp Add a user to db.
func UserSignUp(user *caesar.UserRegistration) error {
	_, err := usersCol.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err.Error())
	}
	return err
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
