package routes

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	productsCollection *mongo.Collection
	client             *mongo.Client
	ctx                context.Context
)

func DatabaseInstance() (context.Context, *mongo.Client) {

	ctx = context.TODO()
	//mongo_uri := os.Getenv("MONGO_URI")
	mongo_uri := "mongodb://localhost:27017"
	conn := options.Client().ApplyURI(mongo_uri)
	client, err := mongo.Connect(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("mongo conn established on product service")
	return ctx, client
}
