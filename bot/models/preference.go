package models

import "go.mongodb.org/mongo-driver/mongo"

type PreferenceModel struct {
	MongoDB *mongo.Database
}

type Preference struct {
	ChatID          int64  `json:"preference_chat_id" bson:"preference_chat_id" `
	EnforcePicture  string `json:"preference_enforce_picture" bson:"preference_enforce_picture" `
	EnforceUsername string `json:"preference_enforce_username" bson:"preference_enforce_username" `
	EnforceAntispam string `json:"preference_enforce_antispam" bson:"preference_enforce_antispam" `
}
