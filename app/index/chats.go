package index

import (
	"github.com/gofiber/fiber/v2"
)

func Chats(c *fiber.Ctx) error {
	return c.Render("chats",
		fiber.Map{
			"Awesome": "Go-Chat",
		})
}
