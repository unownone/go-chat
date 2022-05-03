package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"github.com/unownone/go-chat/routes"
)

func main() {
	err := godotenv.Load()
	if os.Getenv("APP_ENV") == "development" {
		if err != nil {
			fmt.Println("Error loading .env file")
		}
	}
	app := fiber.New(*getConfig())

	// cors
	app.Use(cors.New(*getCorsConfig()))
	// static
	app.Static("/static", "./static")

	// ********************************************
	// 					Routes
	// ********************************************

	//Index Routes
	routes.Index("/", app)
	// Auth Routes
	routes.Auth("/api/v1/auth", app)

	// Chat Routes
	routes.Chat("/api/v1/chat", app)

	// 404 defualt status
	app.Use(get404)

	app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}

func getConfig() *fiber.Config {
	return &fiber.Config{
		Prefork:      true,
		ServerHeader: "Go-Chat-Server",
		AppName:      "Go-Chat",
		Immutable:    true,
		Views:        getHandler(),
	}
}

func getHandler() *html.Engine {
	handler := html.New("./views", ".html")
	handler.Reload(true)
	return handler
}

func getCorsConfig() *cors.Config {
	return &cors.Config{
		AllowCredentials: true,
	}
}

func get404(c *fiber.Ctx) error {
	return c.SendStatus(404)
}
