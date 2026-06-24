package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Database

func ConnectMongo() {
	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("DATABASE_NAME")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("Mongo connection error: ", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("Mongo ping error: ", err)
	}

	DB = client.Database(dbName)

	log.Println("MongoDB connected")

}
