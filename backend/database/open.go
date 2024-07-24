package database

import (
	"context"
	"fmt"
	"runRecord/configuration"
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
	clientOptions := options.Client().ApplyURI(configuration.Connectionstring)

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
