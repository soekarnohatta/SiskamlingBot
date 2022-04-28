package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Picture struct {
	UserID  int64 `json:"user_id" bson:"user_id"`
	ChatID  int64 `json:"chat_id" bson:"chat_id"`
	IsMuted bool  `json:"is_muted" bson:"is_muted"`
}

func NewPicture(userID int64, chatID int64, isMuted bool) *Picture {
	return &Picture{
		UserID:  userID,
		ChatID:  chatID,
		IsMuted: isMuted,
	}
}

func GetPictureByID(db *mongo.Database, Id int64) *Picture {
	var picture *Picture
	dat, err := db.Collection("picture").FindOne(context.TODO(), bson.M{"user_id": Id}).DecodeBytes()
	if err != nil {
		return nil
	}

	err = bson.Unmarshal(dat, &picture)
	if err != nil {
		return nil
	}
	return picture
}

func SavePicture(db *mongo.Database, picture *Picture) {
	_, err := db.Collection("picture").UpdateOne(context.TODO(), bson.M{"user_id": picture.UserID}, bson.D{{Key: "$set", Value: picture}}, options.Update().SetUpsert(true))
	if err != nil {
		log.Print(err.Error())
	}
}

func DeletePictureByID(db *mongo.Database, Id int64) {
	_, err := db.Collection("picture").DeleteOne(context.TODO(), bson.M{"user_id": Id})
	if err != nil {
		log.Print(err.Error())
	}
}
