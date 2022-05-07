package chat

import (
	"strings"

	"github.com/gofiber/websocket/v2"
	"github.com/unownone/go-chat/db"
)

func startChat(h *Hub, c *websocket.Conn, user string, msg string) {
	chatName := strings.TrimPrefix(msg, "!startChat ")
	chat, isChat := db.CheckChat(user, chatName)
	if !isChat {
		if h.running {
			h.broadcast <- []byte(msg)
		} else {
			c.WriteMessage(websocket.TextMessage, []byte("Invalid Chat"))
		}
	} else {
		if h.running {
			h.unregister <- c
		}
		*h = *getCurrHub(chatName)
		if !h.clients[c] {
			h.register <- c
			h.running = true
			c.WriteMessage(websocket.TextMessage, []byte("Welcome to "+chat))
			h.broadcast <- []byte(user + " just joined the chat!")
		} else {
			c.WriteMessage(websocket.TextMessage, []byte("You are already in this chat"))
		}
	}
}
