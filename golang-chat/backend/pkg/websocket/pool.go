package websocket

import (
	"fmt"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			fmt.Println("size of connection pool:", len(p.Clients))

			for client := range p.Clients {
				err := client.Conn.WriteJSON(Message{
					Type: 1,
					Body: "New User Joined...",
				})
				if err != nil {
					fmt.Println("client connection error:", err)
					delete(p.Clients, client)
				}
			}

		case client := <-p.Unregister:
			delete(p.Clients, client)
			fmt.Println("size of connection pool:", len(p.Clients))

			for client := range p.Clients {
				err := client.Conn.WriteJSON(Message{
					Type: 1,
					Body: "User Disconnected...",
				})
				if err != nil {
					fmt.Println("client connection error:", err)
					delete(p.Clients, client)
				}
			}
		case message := <-p.Broadcast:
			fmt.Println("sending message to all clients in the pool")
			for client := range p.Clients {
				err := client.Conn.WriteJSON(message)
				if err != nil {
					fmt.Println("client connection error:", err)
					delete(p.Clients, client)
				}
			}
		}
	}
}
