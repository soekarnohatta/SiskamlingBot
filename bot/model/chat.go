package model

import (
	"context"
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

func GetChatByID(db *mongo.Database, ctx context.Context, Id int) (*Chat, error) {
	var chat *Chat
	dat, err := db.Collection("chat").FindOne(ctx, bson.M{"chat_id": Id}).DecodeBytes()
	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(dat, chat)
	return chat, err
}

func SaveChat(db *mongo.Database, ctx context.Context, chat *Chat) error {
	_, err := db.Collection("chat").UpdateOne(ctx, bson.M{"chat_id": chat.ChatID}, bson.D{{"$set", chat}}, options.Update().SetUpsert(true))
	return err
}

func DeleteChatByID(db *mongo.Database, ctx context.Context, Id int) error {
	_, err := db.Collection("chat").DeleteOne(ctx, bson.M{"chat_id": Id})
	return err
}
