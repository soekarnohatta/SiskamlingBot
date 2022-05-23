package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserModel struct {
	MongoDB *mongo.Database
}

type User struct {
	UserID    int64  `json:"user_id" bson:"user_id" `
	Gban      bool   `json:"user_gban" bson:"user_gban" `
	FirstName string `json:"user_first_name" bson:"user_first_name" `
	LastName  string `json:"user_last_name" bson:"user_last_name" `
	UserName  string `json:"user_username" bson:"user_username" `
}

func (u *UserModel) GetUserById(Id int64) (*User, error) {
	var user *User
	dat, err := u.MongoDB.
		Collection("user").
		FindOne(context.TODO(), bson.M{"user_id": Id}).
		DecodeBytes()
	if err != nil {
		return nil, fmt.Errorf("GetUserByID: failed to retrieve data due to: %w", err)
	}

	err = bson.Unmarshal(dat, &user)
	if err != nil {
		return nil, fmt.Errorf("GetUserByID: failed to unmarshal data due to: %w", err)
	}
	return user, nil
}

func (u *UserModel) SaveUser(user *User) error {
	_, err := u.MongoDB.
		Collection("user").
		UpdateOne(
			context.TODO(),
			bson.M{"user_id": user.UserID},
			bson.D{{Key: "$set", Value: user}},
			options.Update().SetUpsert(true),
		)
	if err != nil {
		return fmt.Errorf("SaveUser: failed to save data due to: %w", err)
	}
	return nil
}

func (u *UserModel) DeleteUserById(Id int64) error {
	_, err := u.MongoDB.Collection("user").DeleteOne(context.TODO(), bson.M{"user_id": Id})
	if err != nil {
		return fmt.Errorf("DeleteUserByID: failed to delete data due to: %w", err)
	}
	return nil
}
