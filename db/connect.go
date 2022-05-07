package db

import (
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoOnce sync.Once
	db_name   string = os.Getenv("MONGO_DB")
	connUri   string = os.Getenv("MONGO_URI")
)

func GetClient() (*mongo.Client, error) {
	var err error
	var client *mongo.Client
	// Using Do once to evalutate the function only once!
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(connUri)

		client, err = mongo.Connect(context.TODO(), clientOptions)

		err2 := client.Ping(context.TODO(), nil)
		if err2 != nil {
			err = err2
		}
	})
	return client, err
}

func GetDatabase() *mongo.Database {
	client, err := GetClient()
	if err != nil {
		panic(err)
	}
	db := client.Database(db_name)
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
