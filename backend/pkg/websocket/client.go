package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type`
	Body string `json:"body"`
	From string `json:""from`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		msgType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := Message{
			Type: msgType,
			Body: string(p),
			From: c.Conn.RemoteAddr().String(),
		}

		c.Pool.Boardcast <- message
		fmt.Println("Message Received: %+v\n", message)
	}
}
