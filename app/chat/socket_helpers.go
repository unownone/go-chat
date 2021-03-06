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
		if h.is_running {
			h.broadcast <- []byte(msg)
		} else {
			c.WriteMessage(websocket.TextMessage, []byte("Invalid Chat"))
		}
	} else {
		old_h := *h
		*h = *getCurrHub(chatName)
		if _, ok := h.clients[c]; !ok {
			if old_h.is_running {
				old_h.unregister <- c
			}
			h.register <- c
			h.is_running = true
			c.WriteMessage(websocket.TextMessage, []byte("Welcome to "+chat))
			h.broadcast <- []byte(user + " just joined the chat!")
		} else {
			h.is_running = true
			c.WriteMessage(websocket.TextMessage, []byte("You are already in this chat"))
		}
	}
}
