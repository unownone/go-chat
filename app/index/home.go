package index

import (
	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	return c.Render("index",
		fiber.Map{
			"Awesome": "Go-Chat",
			"Home":    "active",
		})
}

func Signup(c *fiber.Ctx) error {
	return c.Render("signup",
		fiber.Map{
			"Awesome":  "Go-Chat",
			"Register": "active",
		})
}

func Login(c *fiber.Ctx) error {
	return c.Render("login",
		fiber.Map{
			"Awesome": "Go-Chat",
			"Login":   "active",
		})
}
