package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDatabase() *mongo.Database {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic(err)
	}

	db := client.Database("chat_app")
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
