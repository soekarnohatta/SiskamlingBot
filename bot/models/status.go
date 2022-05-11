package models

import "go.mongodb.org/mongo-driver/mongo"

type StatusModel struct {
	MongoDB *mongo.Database
}

type Status struct {
	ChatId  int64  `json:"status_chat_id" bson:"status_chat_id" `
	UserId  int64  `json:"status_user_id" bson:"status_user_id" `
	Type    string `json:"status_type" bson:"status_type" `
	Action  string `json:"status_action" bson:"status_action" `
	Message string `json:"status_message" bson:"status_message" `
}
