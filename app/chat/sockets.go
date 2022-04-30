package chat

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Hub struct {
	// Registered clients.
	clients map[*websocket.Conn]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *websocket.Conn

	// Unregister requests from clients.
	unregister chan *websocket.Conn
}

func GetSocketUpgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func ChatConnection(c *websocket.Conn) {
	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			c.WriteMessage(mt, []byte("Error Occured!"))
		}
		fmt.Printf("recv: %s", msg)
		if err = c.WriteMessage(mt, msg); err != nil {
			fmt.Println("write:", err)
			break
		}
	}
}
