package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDatabase() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	db := client.Database("chatapp")
	return db
}

func GetUserCol() *mongo.Collection {
	return GetDatabase().Collection("users")
}

func GetChatCol() *mongo.Collection {
	return GetDatabase().Collection("chats")
}

func GetMessageCol() *mongo.Collection {
	return GetDatabase().Collection("messages")
}
