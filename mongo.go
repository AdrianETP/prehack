package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func MongoConnect() {
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(""))
	if err != nil {
		panic(err)
	}

	Database = db.Database("Prehack")
}
