package index

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Chats(c *fiber.Ctx) error {
	return c.Render("chats",
		fiber.Map{
			"HOST": os.Getenv("HOST"),
			"PORT": os.Getenv("PPort"),
		})
}
