package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/unownone/go-chat/app/response"
	"github.com/unownone/go-chat/db"
	"go.mongodb.org/mongo-driver/bson"
)

func CurrentUser(c *fiber.Ctx, claims *jwt.RegisteredClaims) error {
	return c.JSON(
		fiber.Map{
			"user": claims.Issuer,
		},
	)
}

func GetUserByEmail(c *fiber.Ctx) error {
	users := db.GetUserCol()
	email := c.Query("email")
	result := users.FindOne(context.TODO(), bson.M{"email": email})
	if result.Err() != nil {
		return c.JSON(
			response.Error{
				Message: "User not found",
				Error:   true,
			},
		)
	} else {
		user := new(db.User)
		err := result.Decode(user)
		if err != nil {
			return c.JSON(
				response.Error{
					Message: "Error while decoding user",
					Error:   true,
				},
			)
		} else {
			user.Password = ""
			return c.JSON(
				user,
			)
		}
	}
}
