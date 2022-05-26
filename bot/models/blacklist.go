package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlacklistModel struct {
	MongoDB *mongo.Database
}

type Blacklist struct {
	BlacklistTrigger string `json:"blacklist_trigger" bson:"blacklist_trigger" `
}

func (m *BlacklistModel) GetBlacklistByTrigger(Trigger string) (*Blacklist, error) {
	var blacklist *Blacklist
	dat, err := m.MongoDB.
		Collection("blacklist").
		FindOne(context.TODO(), bson.M{"blacklist_trigger": Trigger}).
		DecodeBytes()
	if err != nil {
		return nil, fmt.Errorf("GetBlacklistByTrigger: failed to retrieve data due to: %w", err)
	}

	err = bson.Unmarshal(dat, &blacklist)
	if err != nil {
		return nil, fmt.Errorf("GetBlacklistByTrigger: failed to unmarshal data due to: %w", err)
	}
	return blacklist, nil
}

func (m *BlacklistModel) GetAllBlacklist() ([]Blacklist, error) {
	var allBlacklist []Blacklist
	dat, err := m.MongoDB.
		Collection("blacklist").
		Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("GetBlacklistByTrigger: failed to retrieve data due to: %w", err)
	}

	err = dat.All(context.TODO(), &allBlacklist)
	if err != nil {
		return nil, fmt.Errorf("GetBlacklistByTrigger: failed to parse data due to: %w", err)
	}

	return allBlacklist, nil
}

func (m *BlacklistModel) SaveBlacklist(blacklist *Blacklist) error {
	_, err := m.MongoDB.
		Collection("blacklist").
		UpdateOne(
			context.TODO(),
			bson.M{"blacklist_trigger": blacklist.BlacklistTrigger},
			bson.D{{Key: "$set", Value: blacklist}},
			options.Update().SetUpsert(true),
		)
	if err != nil {
		return fmt.Errorf("SaveBlacklist: failed to save data due to: %w", err)
	}
	return nil
}

func (m *BlacklistModel) DeleteBlacklistByTrigger(Trigger string) error {
	_, err := m.MongoDB.
		Collection("blacklist").
		DeleteOne(context.TODO(), bson.M{"blacklist_trigger": Trigger})
	if err != nil {
		return fmt.Errorf("DeleteBlacklistByTrigger: failed to save data due to: %w", err)
	}
	return nil
}
