package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PreferenceModel struct {
	MongoDB *mongo.Database
}

type Preference struct {
	PreferenceID         int64 `json:"preference_chat_id" bson:"preference_chat_id" `
	EnforcePicture       bool  `json:"preference_enforce_picture" bson:"preference_enforce_picture" `
	EnforceUsername      bool  `json:"preference_enforce_username" bson:"preference_enforce_username" `
	EnforceAntispam      bool  `json:"preference_enforce_antispam" bson:"preference_enforce_antispam" `
	EnforceAntiChinese   bool  `json:"preference_enforce_antichinese" bson:"preference_enforce_antichinese" `
	EnforceAntiArab      bool  `json:"preference_enforce_antiarab" bson:"preference_enforce_antiaran" `
	LastServiceMessageId int64 `json:"preference_last_service_id" bson:"preference_last_service_id" `
}

func (m *PreferenceModel) GetPreferenceById(Id int64) (*Preference, error) {
	var preference *Preference
	dat, err := m.MongoDB.
		Collection("preference").
		FindOne(context.TODO(), bson.M{"preference_chat_id": Id}).
		DecodeBytes()
	if err != nil {
		return nil, fmt.Errorf("GetPreferenceByID: failed to retrieve data due to: %w", err)
	}

	err = bson.Unmarshal(dat, &preference)
	if err != nil {
		return nil, fmt.Errorf("GetPreferenceByID: failed to unmarshal data due to: %w", err)
	}
	return preference, nil
}

func (m *PreferenceModel) SavePreference(preference *Preference) error {
	_, err := m.MongoDB.
		Collection("preference").
		UpdateOne(
			context.TODO(),
			bson.M{"preference_chat_id": preference.PreferenceID},
			bson.D{{Key: "$set", Value: preference}},
			options.Update().SetUpsert(true),
		)
	if err != nil {
		return fmt.Errorf("SavePreference: failed to save data due to: %w", err)
	}
	return nil
}

func (m *PreferenceModel) DeletePreferenceById(Id int64) error {
	_, err := m.MongoDB.
		Collection("preference").
		DeleteOne(context.TODO(), bson.M{"preference_chat_id": Id})
	if err != nil {
		return fmt.Errorf("DeletePreferenceByID: failed to save data due to: %w", err)
	}
	return nil
}
