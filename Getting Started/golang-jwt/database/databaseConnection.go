package database

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func DBinstace() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	MongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

var Client *mongo.Client = DBinstace()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}
