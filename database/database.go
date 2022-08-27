package database

import (
	"context"
	"fmt"
	"time"

	"github.com/OneKiwiTech/gofiber/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	databaseURL  = config.LoadConfig("DB_URL")
	databaseName = config.LoadConfig("DB_NAME")
	DB           *mongo.Database
)

func Connect() error {

	client, err := mongo.NewClient(options.Client().ApplyURI(databaseURL))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB.")
	DB = client.Database(databaseName)
	return nil
}

func Collection(collectionName string) *mongo.Collection {
	var collection *mongo.Collection = DB.Collection(collectionName)
	return collection
}
