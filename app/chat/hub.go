package chat

import (
	"fmt"
	"time"

	"github.com/gofiber/websocket/v2"
)

var (
	hubs    = make(map[string](chan Hub))
	current *websocket.Conn
)

type Hub struct {
	current chan *websocket.Conn
	// Registered clients.
	clients map[*websocket.Conn]bool

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
			h.clients[connection] = true

		case message := <-h.broadcast:
			for connection := range h.clients {
				if current != connection {
					if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						h.unregister <- connection
						connection.WriteMessage(websocket.CloseMessage, []byte{})
						connection.Close()
					}
				}
			}

		case connection := <-h.unregister:
			delete(h.clients, connection)
			if len(h.clients) == 0 {
				h.running = false
			}

		case Client := <-h.current:
			current = Client
		}

	}
}

func newHub() *Hub {
	return &Hub{
		current:    make(chan *websocket.Conn),
		broadcast:  make(chan []byte),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		clients:    make(map[*websocket.Conn]bool),
		running:    true,
	}
}

func getCurrHub(chat string) *Hub {
	if hubs[chat] == nil {
		hubs[chat] = make(chan Hub)

		hubs[chat] <- *newHub()
	}
	// fmt.Println("\nhub: ", <-hubs[chat])
	final := <-hubs[chat]
	fmt.Println("Current: ", final)
	hubs[chat] <- final
	return &final
}
