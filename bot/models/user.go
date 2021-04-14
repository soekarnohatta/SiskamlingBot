package models

import (
	"SiskamlingBot/bot/helpers/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	UserID    int64  `json:"user_id" bson:"user_id" `
	FirstName string `json:"user_first_name" bson:"user_first_name" `
	LastName  string `json:"user_last_name" bson:"user_last_name" `
	UserName  string `json:"user_username" bson:"user_username" `
}

func GetUserByID(ctx context.Context, Id int) (*User, error) {
	var user User
	dat, err := database.Mongo.Collection("user").FindOne(ctx, bson.M{"user_id": Id}).DecodeBytes()
	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(dat, user)
	return &user, err
}

func SaveUser(ctx context.Context, user User) error {
	_, err := database.Mongo.Collection("user").UpdateOne(ctx, bson.M{"user_id": user.UserID}, bson.D{{"$set", user}}, options.Update().SetUpsert(true))
	return err
}

func DeleteUserByID(ctx context.Context, Id int) error {
	_, err := database.Mongo.Collection("user").DeleteOne(ctx, bson.M{"user_id": Id})
	return err
}
