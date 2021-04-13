package database

import (
	"context"
	"github.com/soekarnohatta/Siskamling/bot"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Database

func NewMongo() {
	ctx, cancel := context.WithTimeout(context.TODO(), 10 * time.Second)
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
	log.Println("succesfully Connected to MongoDB Instance!")
}
