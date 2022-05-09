package db

import (
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoOnce  sync.Once
	mongoTwice sync.Once

	db_name string = os.Getenv("MONGO_DB")
	connUri string = os.Getenv("MONGO_URI")
	db      *mongo.Database
	client  *mongo.Client

	user    *mongo.Collection
	chat    *mongo.Collection
	message *mongo.Collection
)

func getClient() (*mongo.Client, error) {
	var err error
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
	mongoTwice.Do(func() {
		client, err := getClient()
		if err != nil {
			panic(err)
		}
		db = client.Database(db_name)
	})
	return db
}

func GetUserCol() *mongo.Collection {
	if user == nil {
		user = GetDatabase().Collection("users")
	}
	return user
}

func GetChatCol() *mongo.Collection {
	if chat == nil {
		chat = GetDatabase().Collection("chats")
	}
	return chat
}

func GetMessageCol() *mongo.Collection {
	if message == nil {
		message = GetDatabase().Collection("messages")
	}
	return message
}
