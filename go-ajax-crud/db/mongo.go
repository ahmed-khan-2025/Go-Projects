package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var BookCollection *mongo.Collection
var UserCollection *mongo.Collection

func InitMongo() {
	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	Client = client
	db := client.Database("booksdb")

	BookCollection = db.Collection("books")
	UserCollection = db.Collection("users")

	log.Println("Mongo Connected")
}