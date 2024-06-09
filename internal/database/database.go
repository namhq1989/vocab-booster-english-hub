package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	mongo *mongo.Database
}

func NewDatabaseClient(url, dbName string) *Database {
	var ctx = context.Background()

	// use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	// send a ping to confirm a successful connection
	if err = client.Database("admin").RunCommand(ctx, bson.M{"ping": 1}).Err(); err != nil {
		panic(err)
	}

	fmt.Printf("⚡️ [mongodb]: connected \n")

	return &Database{
		mongo: client.Database(dbName),
	}
}
