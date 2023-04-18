package app

import (
	"SiskamlingBot/bot/models"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	User      models.UserModel
	Chat      models.ChatModel
	Pref      models.PreferenceModel
	Blacklist models.BlacklistModel
}

func (b *MyApp) newMongo() error {
	log.Println("Connecting to MongoDB instance...")
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	newMongo, err := mongo.NewClient(options.Client().ApplyURI(b.Config.DatabaseURL))
	if err != nil {
		return fmt.Errorf("newMongo: failed to create new client with error: %w", err)
	}

	err = newMongo.Connect(ctx)
	if err != nil {
		return fmt.Errorf("newMongo: failed to connect with error: %w", err)
	}

	mongoDB := newMongo.Database("test")
	err = newMongo.Ping(ctx, nil)
	if err != nil {
		if b.Config.IsDebug {
			log.Printf("Mongo URL is: %s\n", b.Config.DatabaseURL)
		}
		return fmt.Errorf("newMongo: failed to ping new client with error: %w", err)
	}

	b.DB.User.MongoDB = mongoDB
	b.DB.Chat.MongoDB = mongoDB
	b.DB.Pref.MongoDB = mongoDB
	b.DB.Blacklist.MongoDB = mongoDB

	log.Println("Successfully connected to MongoDB instance!")
	return nil
}
