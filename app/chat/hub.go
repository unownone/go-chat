package chat

import (
	"github.com/gofiber/websocket/v2"
)

var (
	curr *websocket.Conn
	hubs *Hubs = getHubRun()
)

type Hubs struct {
	hubs map[string]*Hub
	run  chan *Hub
	stop chan *Hub
}
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
	running    chan bool
	is_running bool `default:"false"`
}

func getHubRun() *Hubs {
	return &Hubs{
		hubs: make(map[string]*Hub),
		run:  make(chan *Hub),
		stop: make(chan *Hub),
	}
}

func HubRunner() {
	for {
		select {
		case hub := <-hubs.run:
			go hub.Run()

		case hub := <-hubs.stop:
			hub.running <- false
		}
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
				h.running <- false
			}

		case curr := <-h.running:
			if !curr {
				return
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
		running:    make(chan bool),
	}
}

func getCurrHub(chat string) *Hub {
	if hub, ok := hubs.hubs[chat]; ok {
		return hub
	} else {
		hub := newHub()
		hubs.hubs[chat] = hub
		hubs.run <- hub
		return hub
	}
}
