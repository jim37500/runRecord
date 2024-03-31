package database

import (
	"context"
	"fmt"
	"runRecord/model"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

func Open() {

	ctx := context.Background()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Replace with your connection string

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		time.Sleep(time.Second)
		Open()
		return
	}

	fmt.Println("Connected to MongoDB!")

	model.AutoMigrate(client)

}
