package model

import (
	"context"
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

func GetUsernameByID(db *mongo.Database, ctx context.Context, Id int64) (*Username, error) {
	var username *Username
	dat, err := db.Collection("username").FindOne(ctx, bson.M{"username_id": Id}).DecodeBytes()
	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(dat, username)
	return username, err
}

func SaveUsername(db *mongo.Database, ctx context.Context, username *Username) error {
	_, err := db.Collection("username").UpdateOne(ctx, bson.M{"username_id": username.UserID}, bson.D{{"$set", username}}, options.Update().SetUpsert(true))
	return err
}

func DeleteUsernameByID(db *mongo.Database, ctx context.Context, Id int64) error {
	_, err := db.Collection("username").DeleteOne(ctx, bson.M{"username_id": Id})
	return err
}
