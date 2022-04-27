package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/unownone/gweeter/app/response"
	"github.com/unownone/gweeter/db"
	"go.mongodb.org/mongo-driver/bson"
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
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return c.JSON(
			response.Error{
				Message: "Invalid Password",
				Error:   true,
			},
		)
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
	_, err = users.InsertOne(context.TODO(), user)
	if err != nil {
		return c.JSON(
			response.Error{
				Message: "Error while creating user",
				Error:   true,
			},
		)
	} else {
		token, err := GetJwtToken(user.Email)
		if err != nil {
			return c.JSON(response.Error{
				Message: "Error while creating token",
				Error:   true,
			})
		}
		return c.JSON(
			response.Success{
				Access:  token,
				Message: "User created successfully",
				Error:   false,
			},
		)
	}
}

func CurrentUser(c *fiber.Ctx) error {
	return c.JSON(
		fiber.Map{
			"message": "User logged in successfully",
		},
	)
}

// func handleLogin(c *fiber.Ctx) error {

// }

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
