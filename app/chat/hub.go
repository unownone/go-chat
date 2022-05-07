package chat

import (
	"time"

	"github.com/gofiber/websocket/v2"
)

var (
	hubs = make(map[string](chan Hub))
	curr *websocket.Conn
)

type Hub struct {
	current chan *websocket.Conn
	// Registered clients.
	clients map[*websocket.Conn]string

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *websocket.Conn

	// Unregister requests from clients.
	unregister chan *websocket.Conn

	// Check if hub is running
	running bool `default:"false"`
}

func HubRunner() {
	for {
		for _, hub := range hubs {
			hub_data := <-hub
			if hub_data.running {
				hub_data.running = false
				go hub_data.Run()
			}
			hub <- hub_data
		}
		time.Sleep(time.Second * 1)
	}
}

func (h *Hub) Run() {
	for {
		select {
		case connection := <-h.register:
			h.clients[connection] = connection.Params("uid")

		case message := <-h.broadcast:
			for connection := range h.clients {
				if curr != connection {
					if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						h.unregister <- connection
						connection.WriteMessage(websocket.CloseMessage, []byte("Error Occured!"))
						connection.Close()
					}
				}
			}

		case connection := <-h.unregister:
			delete(h.clients, connection)
			if len(h.clients) == 0 {
				h.running = false
			}

		case connection := <-h.current:
			curr = connection

		}
	}
}

func newHub() *Hub {
	return &Hub{
		current:    make(chan *websocket.Conn),
		broadcast:  make(chan []byte),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		clients:    make(map[*websocket.Conn]string),
		running:    true,
	}
}

func getCurrHub(chat string) *Hub {
	var final Hub
	if hubs[chat] == nil {
		hubs[chat] = make(chan Hub)
		final = *newHub()
		hubs[chat] <- final
	} else {
		final = <-hubs[chat]
		hubs[chat] <- final
	}
	return &final
}
