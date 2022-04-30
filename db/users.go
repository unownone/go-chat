package db

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

func GetCurrUser(claims *jwt.RegisteredClaims, user *User) error {
	users := GetUserCol()
	user_email := claims.Issuer
	err := users.FindOne(context.TODO(), bson.M{"email": user_email}).Decode(user)
	return err
}
