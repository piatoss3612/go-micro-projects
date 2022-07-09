package main

import (
	"fmt"
	"golang-chat/pkg/websocket"
	"log"
	"net/http"
)

// upgrade connection to websocket and register client to websocket pool
func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("websocket endpoint reached")

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		log.Panic(err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start() // start goroutine that listens on channels in the websocket pool

	// request to localhost:9000/ws
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func main() {
	fmt.Println("starting backend...")

	setupRoutes()

	// start web server on localhost:9000
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Panic(err)
	}
}
