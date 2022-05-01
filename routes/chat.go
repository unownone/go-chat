package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/unownone/go-chat/app/chat"
	"github.com/unownone/go-chat/middleware"
)

// Base Route = /api/v1/chat
func Chat(base string, app *fiber.App) {

	// Get Chats
	app.Get(base+"/chats", middleware.VerifyJwtWithClaim(chat.GetChats))
	// Create Chat
	app.Post(base+"/chats", middleware.VerifyJwtWithClaim(chat.CreateChat))
	go chat.HubRunner()
	app.Use(base, chat.GetSocketUpgrade)
	app.Get(base+"/:sess/:id", websocket.New(chat.ChatConnection))
}
