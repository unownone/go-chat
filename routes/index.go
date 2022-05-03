package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unownone/go-chat/app/index"
)

// Base Route = /
func Index(base string, app *fiber.App) {
	app.Get(base, index.Homepage)
	app.Get(base+"signup", index.Signup)
	app.Get(base+"login", index.Login)
	app.Get(base+"chats", index.Chats)
}
