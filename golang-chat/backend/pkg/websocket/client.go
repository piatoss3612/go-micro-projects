package websocket

import (
	"fmt"
	"log"
	_ "sync"

	"github.com/gorilla/websocket"
)

// websocket client
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
	// mu   sync.Mutex
}

// websocket payload
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

// read message sent from clients and broadcast it in infinite loop
// if something went wrong while reading messages,
// defered literal function will close websocket connection
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := Message{
			Type: messageType,
			Body: string(p),
		}

		c.Pool.Broadcast <- message
		fmt.Printf("message received: %v\n", message)
	}
}
