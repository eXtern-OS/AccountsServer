package db

import (
	"context"
	beatrix "github.com/eXtern-OS/Beatrix"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var URI = ""

func Init(mongouri string) {
	URI = mongouri
}

func NewDBCollection(collectionName, issuer string) (bool, *mongo.Collection) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Println(err)
		go beatrix.SendError("Error creating client for mongodb", issuer)
		return false, nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		go beatrix.SendError("Error connecting with client to mongodb", issuer)
		return false, nil
	}

	collection := client.Database("Users").Collection(collectionName)
	return true, collection
}

func NewDatabaseCollection(database, collectionName, issuer string) (bool, *mongo.Collection) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Println(err)
		go beatrix.SendError("Error creating client for mongodb", issuer)
		return false, nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		go beatrix.SendError("Error connecting with client to mongodb", issuer)
		return false, nil
	}

	collection := client.Database(database).Collection(collectionName)
	return true, collection
}
