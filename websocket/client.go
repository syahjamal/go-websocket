package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

// Notification struct
type Notification struct {
	gorm.Model
	ID       int `json:"id"`
	UserIDTo string
	Username string
	Message  string `json:"message"`
}

// Client struct
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

func (c *Client) Read() {
	defer func() {
		// c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := Notification{ID: messageType, Message: string(p)}
		fmt.Printf("Message Received: %+v\n", message)

	}
}
