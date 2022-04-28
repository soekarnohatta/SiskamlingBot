package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Username struct {
	UserID  int64 `json:"user_id" bson:"user_id"`
	ChatID  int64 `json:"chat_id" bson:"chat_id"`
	IsMuted bool  `json:"is_muted" bson:"is_muted"`
}

func NewUsername(userID int64, chatID int64, isMuted bool) *Username {
	return &Username{
		UserID:  userID,
		ChatID:  chatID,
		IsMuted: isMuted,
	}
}

func GetUsernameByID(db *mongo.Database, Id int64) *Username {
	var username *Username
	dat, err := db.Collection("username").FindOne(context.TODO(), bson.M{"user_id": Id}).DecodeBytes()
	if err != nil {
		return nil
	}

	err = bson.Unmarshal(dat, &username)
	if err != nil {
		return nil
	}
	return username
}

func SaveUsername(db *mongo.Database, username *Username) {
	_, err := db.Collection("username").UpdateOne(context.TODO(), bson.M{"user_id": username.UserID}, bson.D{{Key: "$set", Value: username}}, options.Update().SetUpsert(true))
	if err != nil {
		log.Print(err.Error())
	}
	return
}

func DeleteUsernameByID(db *mongo.Database, Id int64) {
	_, err := db.Collection("username").DeleteOne(context.TODO(), bson.M{"user_id": Id})
	if err != nil {
		log.Print(err.Error())
	}
	return
}
