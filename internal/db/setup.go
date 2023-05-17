package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"modelo-rest-go/internal/configs"
	"time"
)

var Client *mongo.Client

func ConnectDB() {

	conf := configs.GetConfig()
	uriMongo := conf.GetString("db.mongouri")

	client, err := mongo.NewClient(options.Client().ApplyURI(uriMongo))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Mongo")

	Client = client
}

func GetCollection(client *mongo.Client, database string, collectionName string) *mongo.Collection {
	collection := client.Database(database).Collection(collectionName)
	return collection
}
