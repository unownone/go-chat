package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/unownone/go-chat/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Auth Routes
	routes.Auth("/api/v1/auth", app)

	// Chat Routes
	routes.Chat("/api/v1/chat", app)
	app.Listen("localhost:" + os.Getenv("PORT"))
}
