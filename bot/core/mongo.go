package core

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (b *MyApp) newMongo() {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	newMongo, err := mongo.NewClient(options.Client().ApplyURI(b.Config.DatabaseURL))
	if err != nil {
		log.Fatalln("cannot create mongo client: " + err.Error())
	}

	err = newMongo.Connect(ctx)
	if err != nil {
		panic("cannot connect to mongo database: " + err.Error())
	}

	b.DB = newMongo.Database("test")

	// Ping check to minimize error during long run.
	err = newMongo.Ping(ctx, nil)
	if err != nil {
		if b.Config.IsDebug {
			log.Printf("mongo url is: %s\n", b.Config.DatabaseURL)
		}
		log.Fatalln("cannot connect to mongo database: " + err.Error())
	}

	log.Println("successfully connected to MongoDB instance!")
}
