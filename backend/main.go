package main

import (
	"fmt"
	"log"
	"net/http"

	"chat-app/pkg/websocket"
)

// WebSocket endpoint
func serverWs(p *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket EndPoint Hit")
	// upgrade to websocket
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: p,
	}

	p.Register <- client

	client.Read()
}

func setupRoutes() {
	p := websocket.NewPool()
	go p.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWs(p, w, r)
	})
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
