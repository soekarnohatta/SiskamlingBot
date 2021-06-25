package app

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (b *MyApp) newMongo() {
	log.Print("Connecting to MongoDB instance...")
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	newMongo, err := mongo.NewClient(options.Client().ApplyURI(b.Config.DatabaseURL))
	if err != nil {
		log.Fatal("Cannot create mongo client: " + err.Error())
	}

	err = newMongo.Connect(ctx)
	if err != nil {
		log.Fatal("Cannot connect to mongo database: " + err.Error())
	}

	b.DB = newMongo.Database("test")

	err = newMongo.Ping(ctx, nil)
	if err != nil {
		if b.Config.IsDebug {
			log.Printf("Mongo URL is: %s\n", b.Config.DatabaseURL)
		}
		log.Fatal("Cannot connect to mongo database: " + err.Error())
	}

	log.Print("Successfully connected to MongoDB instance!")
}
