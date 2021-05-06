package database

import (
	"SiskamlingBot/bot"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Database

func NewMongo() {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	newMongo, err := mongo.NewClient(options.Client().ApplyURI(bot.Config.DatabaseURL))
	if err != nil {
		panic("cannot create mongo client: " + err.Error())
	}

	err = newMongo.Connect(ctx)
	if err != nil {
		panic("cannot connect to mongo database: " + err.Error())
	}

	Mongo = newMongo.Database("test")

	// Ping check to minimize error during long run.
	err = newMongo.Ping(ctx, nil)
	if err != nil {
		if bot.Config.IsDebug {
			log.Printf("mongo url is: %s", bot.Config.DatabaseURL)
		}
		panic("cannot connect to mongo database: " + err.Error())
	}

	log.Println("successfully connected to MongoDB instance!")
}
