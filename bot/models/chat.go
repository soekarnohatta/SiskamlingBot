package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Chat struct {
	ChatID    int64  `json:"chat_id" bson:"chat_id" `
	ChatType  string `json:"chat_type" bson:"chat_type" `
	ChatLink  string `json:"chat_link" bson:"chat_link" `
	ChatTitle string `json:"chat_title" bson:"chat_title" `
}

func NewChat(ID int64, chatType string, chatLink string, chatTitle string) *Chat {
	return &Chat{
		ChatID:    ID,
		ChatType:  chatType,
		ChatLink:  chatLink,
		ChatTitle: chatTitle,
	}
}

func GetChatByID(db *mongo.Database, Id int64) *Chat {
	var chat *Chat
	dat, err := db.Collection("chat").FindOne(context.TODO(), bson.M{"chat_id": Id}).DecodeBytes()
	if err != nil {
		return nil
	}

	err = bson.Unmarshal(dat, &chat)
	if err != nil {
		return nil
	}
	return chat
}

func SaveChat(db *mongo.Database, chat *Chat) {
	_, err := db.Collection("chat").UpdateOne(context.TODO(), bson.M{"chat_id": chat.ChatID}, bson.D{{Key: "$set", Value: chat}}, options.Update().SetUpsert(true))
	if err != nil {
		log.Print(err.Error())
	}
	return
}

func DeleteChatByID(db *mongo.Database, Id int64) {
	_, err := db.Collection("chat").DeleteOne(context.TODO(), bson.M{"chat_id": Id})
	if err != nil {
		log.Print(err.Error())
	}
	return
}
