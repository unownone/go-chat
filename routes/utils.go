package routes

import "github.com/gofiber/fiber/v2"

func GetNextMiddleWare(c *fiber.Ctx) error {
	return c.Next()
}
