package websocket

import (
	"fmt"
)

// websocket pool
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

// factory function
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

// start listening on websocket pool
func (p *Pool) Start() {
	for {
		select {
		// register client
		case client := <-p.Register:
			p.Clients[client] = true
			fmt.Println("size of connection pool:", len(p.Clients))

			msg := Message{
				Type: 1,
				Body: "New User Joined...",
			}

			p.BroadcastMessage(msg)

		// unregister client
		case client := <-p.Unregister:
			delete(p.Clients, client)
			fmt.Println("size of connection pool:", len(p.Clients))

			msg := Message{
				Type: 1,
				Body: "User Disconnected...",
			}

			p.BroadcastMessage(msg)

		// broadcast chat message
		case msg := <-p.Broadcast:
			fmt.Println("sending message to all clients in the pool")

			p.BroadcastMessage(msg)
		}
	}
}

// separate broadcasting message
func (p *Pool) BroadcastMessage(msg Message) {
	for client := range p.Clients {
		err := client.Conn.WriteJSON(msg)
		if err != nil {
			fmt.Println("client connection error:", err)
			delete(p.Clients, client)
		}
	}
}
