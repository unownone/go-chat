package db

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID     primitive.ObjectID   `bson:"_id" json:"id,"`
	Name   string               `bson:"name,omitempty" json:"name"`
	Users  []primitive.ObjectID `bson:"users,omitempty" json:"users"`
	Public bool                 `bson:"public,omitempty" json:"public" default:"false"`
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

func UpdateChat(chat *Chat) error {
	chat_col := GetChatCol()
	another := new(Chat)
	err := chat_col.FindOne(context.TODO(), bson.M{"_id": chat.ID}).Decode(another)
	if err != nil {
		return err
	}
	getFinalChat(chat, another)
	_, err = chat_col.ReplaceOne(context.TODO(), bson.M{"_id": chat.ID}, another)
	return err
}

// Helper function used to get a updated chat , removing users that are already there and adding new valid users
//
// It reads data from chat and returns a updated another chat that would be used to update chat function
func getFinalChat(chat *Chat, another *Chat) {
	if chat.Name != "" {
		another.Name = chat.Name
	}
	for _, user := range chat.Users {
		if !contains(another.Users, user) {
			if userIsValid(user) {
				another.Users = append(chat.Users, user)
			}
		}
	}
}

func userIsValid(user primitive.ObjectID) bool {
	users := new(User)
	err := GetUserById(user, users)
	return err == nil
}

func contains(s []primitive.ObjectID, e primitive.ObjectID) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Checks whether User has access to Chat
func CheckChat(user string, chat string) (string, bool) {
	users := new(User)
	claims := jwt.RegisteredClaims{Issuer: user}
	err := GetCurrUser(&claims, users)
	if err != nil {
		return "", false
	}
	chats := new([]Chat)
	hexchat, err := primitive.ObjectIDFromHex(chat)
	if err != nil {
		return users.Name, false
	}
	err = GetUserChats(users.ID, chats)
	if err != nil {
		return users.Name, false
	}
	for _, c := range *chats {
		if c.ID == hexchat {
			return c.Name, true
		}
	}
	return users.Name, false
}
