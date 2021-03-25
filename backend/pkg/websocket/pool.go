package websocket

import (
	"fmt"
	"log"
)

type Pool struct {
	Register   chan *Client // New User Joined
	Unregister chan *Client // unregister a user
	Clients    map[*Client]bool
	Boardcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Boardcast:  make(chan Message),
	}
}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(p.Clients))
			for client, _ := range p.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{
					Type: 1,
					Body: "New User Joined...",
				})
			}
		case client := <-p.Unregister:
			delete(p.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(p.Clients))
			for client, _ := range p.Clients {
				client.Conn.WriteJSON(Message{
					Type: 1,
					Body: "User Disconnected...",
				})
			}
		case message := <-p.Boardcast:
			fmt.Println("Sending message to all clients in Pool")
			for client, _ := range p.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}
