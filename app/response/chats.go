package response

import "github.com/unownone/go-chat/db"

type ChatResponse struct {
	Message string    `json:"message"`
	Error   bool      `json:"error" default:"false"`
	Chats   []db.Chat `json:"chats"`
}
