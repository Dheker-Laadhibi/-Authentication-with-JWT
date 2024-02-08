package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go/mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("err loading .env file")
	}
	Mongodb := os.Getenv("MONGODB_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(Mongodb))
	if err != nil {

		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		fmt.Println("connected to mongo db")
	}
	return client
}

var Client *mongo.Client = DBinstance()

func openCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var Collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return Collection
}
