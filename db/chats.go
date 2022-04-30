package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID    primitive.ObjectID   `bson:"_id" json:"id,"`
	Name  string               `bson:"name,omitempty" json:"name"`
	Users []primitive.ObjectID `bson:"users,omitempty" json:"users"`
}

type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ChatID     primitive.ObjectID `bson:"chat_id,omitempty" json:"chat_id"`
	FromUserID primitive.ObjectID `bson:"from_user_id,omitempty" json:"from_user_id"`
	Message    string             `bson:"message,omitempty" json:"message"`
	Time       primitive.DateTime `bson:"time,omitempty" json:"time"`
}

// Helper Functions

func GetUserChats(id primitive.ObjectID, chats *[]Chat) error {
	chat_col := GetChatCol()
	filter := bson.M{"users": bson.M{"$elemMatch": bson.M{"$eq": id}}}
	cursor, err := chat_col.Find(context.TODO(), filter)
	if err != nil {
		return err
	} else {
		cursor.All(context.TODO(), chats)
		return nil
	}
}

func CreateChat(chat *Chat) error {
	chat.ID = primitive.NewObjectID()
	chat_col := GetChatCol()
	_, err := chat_col.InsertOne(context.TODO(), chat)
	return err
}
