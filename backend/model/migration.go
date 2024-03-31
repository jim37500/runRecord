package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

func AutoMigrate(client *mongo.Client) {
	createRunRecordCollection(client, "run_records")
}

func createRunRecordCollection(client *mongo.Client, collectionName string) {

	db := client.Database("runRecord")
	command := bson.D{{Key: "create", Value: collectionName}}
	var result bson.M
	_ = db.RunCommand(context.TODO(), command).Decode(&result)
}
