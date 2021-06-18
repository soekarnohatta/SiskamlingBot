package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	UserID    int64  `json:"user_id" bson:"user_id" `
	FirstName string `json:"user_first_name" bson:"user_first_name" `
	LastName  string `json:"user_last_name" bson:"user_last_name" `
	UserName  string `json:"user_username" bson:"user_username" `
}

func NewUser(userID int64, firstName string, lastName string, userName string) *User {
	return &User{
		UserID:    userID,
		FirstName: firstName,
		LastName:  lastName,
		UserName:  userName,
	}
}

func GetUserByID(db *mongo.Database, ctx context.Context, Id int) (*User, error) {
	var user *User
	dat, err := db.Collection("user").FindOne(ctx, bson.M{"user_id": Id}).DecodeBytes()
	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(dat, &user)
	return user, err
}

func SaveUser(db *mongo.Database, ctx context.Context, user *User) error {
	_, err := db.Collection("user").UpdateOne(ctx, bson.M{"user_id": user.UserID}, bson.D{{Key: "$set", Value: user}}, options.Update().SetUpsert(true))
	return err
}

func DeleteUserByID(db *mongo.Database, ctx context.Context, Id int) error {
	_, err := db.Collection("user").DeleteOne(ctx, bson.M{"user_id": Id})
	return err
}
