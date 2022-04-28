package app

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (b *MyApp) newMongo() error {
	log.Println("Connecting to MongoDB instance...")
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	newMongo, err := mongo.NewClient(options.Client().ApplyURI(b.Config.DatabaseURL))
	if err != nil {
		return err
	}

	err = newMongo.Connect(ctx)
	if err != nil {
		return err
	}

	b.DB = newMongo.Database("test")

	err = newMongo.Ping(ctx, nil)
	if err != nil {
		if b.Config.IsDebug {
			log.Printf("Mongo URL is: %s\n", b.Config.DatabaseURL)
		}
		return err
	}

	log.Println("Successfully connected to MongoDB instance!")
	return nil
}
