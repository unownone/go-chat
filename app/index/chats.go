package index

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Chats(c *fiber.Ctx) error {
	Env := false
	if os.Getenv("APP_ENV") == "development" {
		Env = true
	}
	return c.Render("chats",
		fiber.Map{
			"HOST":  os.Getenv("HHost"),
			"PORT":  os.Getenv("PPort"),
			"Chats": "active",
			"Env":   Env,
		})
}
