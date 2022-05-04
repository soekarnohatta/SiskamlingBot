package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChatModel struct {
	MongoDB *mongo.Database
}

type Chat struct {
	ChatID    int64  `json:"chat_id" bson:"chat_id" `
	ChatType  string `json:"chat_type" bson:"chat_type" `
	ChatLink  string `json:"chat_link" bson:"chat_link" `
	ChatTitle string `json:"chat_title" bson:"chat_title" `
}

func NewChat(ID int64, chatType, chatLink, chatTitle string) *Chat {
	return &Chat{
		ChatID:    ID,
		ChatType:  chatType,
		ChatLink:  chatLink,
		ChatTitle: chatTitle,
	}
}

func (m *ChatModel) GetChatById(Id int64) (*Chat, error) {
	var chat *Chat
	dat, err := m.MongoDB.Collection("chat").FindOne(context.TODO(), bson.M{"chat_id": Id}).DecodeBytes()
	if err != nil {
		return nil, fmt.Errorf("GetChatByID: failed to retrieve data due to: %w", err)
	}

	err = bson.Unmarshal(dat, &chat)
	if err != nil {
		return nil, fmt.Errorf("GetChatByID: failed to unmarshal data due to: %w", err)
	}
	return chat, nil
}

func (m *ChatModel) SaveChat(chat *Chat) error {
	_, err := m.MongoDB.Collection("chat").UpdateOne(context.TODO(), bson.M{"chat_id": chat.ChatID}, bson.D{{Key: "$set", Value: chat}}, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("SaveChat: failed to save data due to: %w", err)
	}
	return nil
}

func (m *ChatModel) DeleteChatById(Id int64) error {
	_, err := m.MongoDB.Collection("chat").DeleteOne(context.TODO(), bson.M{"chat_id": Id})
	if err != nil {
		return fmt.Errorf("DeleteChatByID: failed to save data due to: %w", err)
	}
	return nil
}
