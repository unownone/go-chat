package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/unownone/go-chat/app/chat"
	"github.com/unownone/go-chat/middleware"
)

// Base Route = /api/:ver/chat
func Chat(base string, app fiber.Router) {
	// Versions
	v1 := app.Group("/v1", GetNextMiddleWare)
	v2 := app.Group("/v2", GetNextMiddleWare)

	// Get Chats
	v1.Get(base+"/chats", middleware.VerifyJwtWithClaim(chat.GetChats))
	// Create Chat
	v1.Post(base+"/chats", middleware.VerifyJwtWithClaim(chat.CreateChat))
	// Update Chat
	v1.Put(base+"/chats", middleware.VerifyJwtWithClaim(chat.UpdateChat))

	v1.Use(base, chat.GetSocketUpgrade)
	v1.Get(base+"/:sess/:id", websocket.New(chat.ChatConnection))

	// Get User Connection and handle logic
	v2.Get(base+"/:sess", websocket.New(chat.AuthWebSocket))
}
