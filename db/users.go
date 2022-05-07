package db

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name,omitempty" json:"name"`
	Email    string             `bson:"email,omitempty" json:"email"`
	Password string             `bson:"password,omitempty" json:"-"`
}

func GetCurrUser(claims *jwt.RegisteredClaims, user *User) error {
	users := GetUserCol()
	user_email := claims.Issuer
	err := users.FindOne(context.TODO(), bson.M{"email": user_email}).Decode(user)
	return err
}

func GetUserById(id primitive.ObjectID, user *User) error {
	users := GetUserCol()
	err := users.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	return err
}
