package chat

import (
	"fmt"
	"strings"

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

// Will be phased out in a future release
//
// Please use chat v2 for better websocket chats.
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
			// fmt.Println("Current", curr_hub)
			(*curr_hub).register <- c

			c.WriteMessage(websocket.TextMessage, []byte("Welcome "+user))
			// fmt.Println("Welcome " + user)

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

// Authorizing Websockets , detaching from the chat selection
// Makign it more efficient
func AuthWebSocket(c *websocket.Conn) {
	var authorized = false
	var user string = ""
	sess := c.Params("sess")
	if sess == "" {
		c.WriteMessage(websocket.TextMessage, []byte("Unauthorized"))
		c.Close()
		return
	}
	user, authorized = auth.VerifyJwtSess(sess)
	if !authorized {
		c.WriteMessage(websocket.TextMessage, []byte("Unauthorized"))
		c.Close()
		return
	}
	user_obj := db.User{}
	db.GetUserByEmail(user, &user_obj)
	c.WriteMessage(websocket.TextMessage, []byte("Welcome "+user_obj.Name))
	// Close if not Authorised
	if !authorized {
		return
	}

	webSocketHandler(c, &user_obj)
}

func webSocketHandler(c *websocket.Conn, user *db.User) {
	//Close if disconnected
	h := Hub{}
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			h.unregister <- c
			c.Close()
			return
		}
		switch {
		case mt == websocket.TextMessage && strings.HasPrefix(string(msg), "!startChat"):
			startChat(&h, c, user.Email, string(msg))
		case mt == websocket.TextMessage:
			if h.is_running {
				h.current <- c
				h.broadcast <- append([]byte(user.Name+": "), msg[:]...)
			} else {
				c.WriteMessage(websocket.TextMessage, []byte("please join a chat first"))
			}
		}
	}
}
