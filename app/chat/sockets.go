package chat

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	auth "github.com/unownone/go-chat/app/user"
	"github.com/unownone/go-chat/db"
)

func GetSocketUpgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func ChatConnection(c *websocket.Conn) {
	curr_hub := new(Hub)
	Username := ""
	var authorized = false
	sess := c.Params("sess")
	if sess == "" {
		c.WriteMessage(websocket.TextMessage, []byte("UnAuthorized 1"))
		c.Close()
	} else {
		user, verif := auth.VerifyJwtSess(sess)
		if !verif {
			c.WriteMessage(websocket.TextMessage, []byte("UnAuthorized 2"))
			c.Close()
		}
		chatSession := c.Params("id")
		abc, isChat := db.CheckChat(user, chatSession)
		Username = abc
		if isChat {
			authorized = true

			curr_hub = getCurrHub(chatSession)
			fmt.Println("Current: ", *curr_hub)
			(*curr_hub).register <- c
			count := getTotalocc(c, curr_hub.clients)
			if count > 1 {
				Username = Username + strconv.Itoa(count)
			}
			c.WriteMessage(websocket.TextMessage, []byte("Welcome "+user))

		} else {
			c.WriteMessage(websocket.TextMessage, []byte("UnAuthorized 3"))
			c.Close()
		}
	}

	if authorized {
		defer func(h *Hub) {
			h.unregister <- c
			c.Close()
		}(curr_hub)

		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				return
			}
			if mt == websocket.TextMessage {
				curr_hub.current <- c
				(*curr_hub).broadcast <- append([]byte(Username+": "), msg[:]...)

			} else {
				fmt.Println("Message type: ", mt)
			}
		}
	}
}

func getTotalocc(c *websocket.Conn, arr map[*websocket.Conn]bool) int {
	count := 0
	for v := range arr {
		if v == c {
			count++
		}
	}
	return count
}
