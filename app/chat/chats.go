package chat

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/unownone/go-chat/app/response"
	"github.com/unownone/go-chat/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetChats(c *fiber.Ctx, claims *jwt.RegisteredClaims) error {
	curr_user := db.User{}
	chats := []db.Chat{}
	err := db.GetCurrUser(claims, &curr_user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return c.Status(404).JSON(response.Error{Message: "User not found"})
		}
	}
	err = db.GetUserChats(curr_user.ID, &chats)
	if err != nil {
		return c.Status(404).JSON(response.Error{Message: "Chats not Found"})
	}
	return c.JSON(response.ChatResponse{
		Message: strconv.Itoa(len(chats)) + " Chats found",
		Chats:   chats,
	})
}

func CreateChat(c *fiber.Ctx, claims *jwt.RegisteredClaims) error {
	curr_user := db.User{}
	err := db.GetCurrUser(claims, &curr_user)
	if err != nil {
		return c.Status(404).JSON(response.Error{Message: "User not found"})
	}
	chat := new(db.Chat)
	err = c.BodyParser(&chat)
	if err != nil {
		return c.Status(400).JSON(response.Error{Message: "Invalid request body"})
	}
	if chat.Users == nil {
		chat.Users = append(chat.Users, curr_user.ID)
	}
	err = db.CreateChat(chat)
	if err != nil {
		return c.Status(400).JSON(response.Error{Message: "Chat not created"})
	}
	return c.Status(201).JSON(response.ChatResponse{
		Message: "Chat created",
		Chats:   []db.Chat{*chat},
	})
}

func UpdateChat(c *fiber.Ctx, claims *jwt.RegisteredClaims) error {
	curr_user := db.User{}
	err := db.GetCurrUser(claims, &curr_user)
	if err != nil {
		return c.Status(404).JSON(response.Error{Message: "User not found"})
	}
	chat := new(db.Chat)
	err = c.BodyParser(&chat)
	if err != nil {
		return c.Status(400).JSON(response.Error{Message: "Invalid request body"})
	}
	if chat.Users == nil {
		chat.Users = append(chat.Users, curr_user.ID)
	}
	err = db.UpdateChat(chat)
	if err != nil {
		return c.Status(400).JSON(response.Error{Message: "Chat not Updated"})
	}
	return c.Status(201).JSON(response.ChatResponse{
		Message: "Updated",
		Chats:   []db.Chat{*chat},
	})
}
