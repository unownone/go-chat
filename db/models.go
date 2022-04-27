package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

type Chat struct {
	ID    primitive.ObjectID   `bson:"_id,omitempty"`
	Name  string               `bson:"name,omitempty"`
	Users []primitive.ObjectID `bson:"users,omitempty"`
}

type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ChatID     primitive.ObjectID `bson:"chat_id,omitempty"`
	FromUserID primitive.ObjectID `bson:"from_user_id,omitempty"`
	Message    string             `bson:"message,omitempty"`
	Time       primitive.DateTime `bson:"time,omitempty"`
}
