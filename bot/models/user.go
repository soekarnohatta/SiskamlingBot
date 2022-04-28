package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	UserID    int64  `json:"user_id" bson:"user_id" `
	Gban      bool   `json:"user_gban" bson:"user_gban" `
	FirstName string `json:"user_first_name" bson:"user_first_name" `
	LastName  string `json:"user_last_name" bson:"user_last_name" `
	UserName  string `json:"user_username" bson:"user_username" `
}

func NewUser(userID int64, firstName string, lastName string, userName string) *User {
	return &User{
		UserID:    userID,
		Gban:      false,
		FirstName: firstName,
		LastName:  lastName,
		UserName:  userName,
	}
}

func GetUserByID(db *mongo.Database, Id int) *User {
	var user *User
	dat, err := db.Collection("user").FindOne(context.TODO(), bson.M{"user_id": Id}).DecodeBytes()
	if err != nil {
		return nil
	}

	err = bson.Unmarshal(dat, &user)
	if err != nil {
		return nil
	}
	return user
}

func SaveUser(db *mongo.Database, user *User) {
	_, err := db.Collection("user").UpdateOne(context.TODO(), bson.M{"user_id": user.UserID}, bson.D{{Key: "$set", Value: user}}, options.Update().SetUpsert(true))
	if err != nil {
		log.Print(err.Error())
	}
	return
}

func DeleteUserByID(db *mongo.Database, Id int) {
	_, err := db.Collection("user").DeleteOne(context.TODO(), bson.M{"user_id": Id})
	if err != nil {
		log.Print(err.Error())
	}
	return
}
