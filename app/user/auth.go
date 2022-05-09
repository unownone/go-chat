package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/unownone/go-chat/app/response"
	"github.com/unownone/go-chat/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	users := db.GetUserCol()
	user := new(db.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.JSON(response.Error{
			Message: "Invalid Data",
			Error:   true,
		})
	}
	result := users.FindOne(context.TODO(), bson.M{"email": user.Email})
	if result.Err() == nil {
		return c.JSON(
			response.Error{
				Message: "Email ID already exists.",
				Error:   true,
			},
		)
	}
	user.ID = primitive.NewObjectID()
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return c.JSON(
			response.Error{
				Message: "Invalid Password",
				Error:   true,
			},
		)
	}
	_, err = users.InsertOne(context.TODO(), user)
	if err != nil {
		return c.JSON(
			response.Error{
				Message: "Error while creating user",
				Error:   true,
			},
		)
	} else {
		token, refresh, err := GetJwtToken(user.Email)
		if err != 1 {
			return c.JSON(response.Error{
				Message: "Error while creating token",
				Error:   true,
			})
		}
		return c.JSON(
			response.Success{
				Access:  token,
				Refresh: refresh,
				Message: "User created successfully",
				Error:   false,
			},
		)
	}
}

func Login(c *fiber.Ctx) error {
	users := db.GetUserCol()
	user := new(db.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.JSON(response.Error{
			Message: "Invalid Data",
			Error:   true,
		})
	}
	result := users.FindOne(context.TODO(), bson.M{"email": user.Email})
	if result.Err() != nil {
		return c.JSON(
			response.Error{
				Message: "Email ID Doesn't Exist. Please Login",
				Error:   true,
			},
		)
	}
	user_obj := new(db.User)
	result.Decode(&user_obj)
	verified := verifyPassword(user.Password, user_obj.Password)
	if !verified {
		return c.JSON(
			response.Error{
				Message: "Invalid Password",
				Error:   true,
			},
		)
	} else {
		token, refresh, err := GetJwtToken(user.Email)
		if err != 1 {
			return c.JSON(response.Error{
				Message: "Error while creating token",
				Error:   true,
			})
		}
		return c.JSON(
			response.Success{
				Access:  token,
				Refresh: refresh,
				Message: "Logged In Successfully",
				Error:   false,
			},
		)
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
