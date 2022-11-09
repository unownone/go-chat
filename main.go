package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/template/html"
	"github.com/unownone/go-chat/app/chat"
	"github.com/unownone/go-chat/routes"
)

func main() {
	go chat.HubRunner()

	app := fiber.New(*getConfig())

	app.Use(pprof.New())

	api := app.Group("/api", routes.GetNextMiddleWare)
	// cors & logging
	app.Use(cors.New(*getCorsConfig()))
	// static
	app.Static("/static", "./static")

	// ********************************************
	// 					Routes
	// ********************************************

	//Index Routes HTML
	routes.Index("/", app)

	// Auth Routes
	routes.Auth("/auth", api)

	// Chat Routes
	routes.Chat("/chat", api)

	// 404 defualt status
	app.Use(get404)

	app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}

func getConfig() *fiber.Config {
	return &fiber.Config{
		Prefork:      false, // Disable Prefork as need only 1 instance running
		ServerHeader: "iMon",
		AppName:      "Go-Chat",
		Immutable:    true,
		Views:        getHandler(),
	}
}

func getHandler() *html.Engine {
	handler := html.New("./views", ".html")
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
