package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

// ReadPump handles incoming messages from the WebSocket connection.
func (c *Client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		c.hub.broadcast <- message
	}
}

// WritePump handles sending messages to the WebSocket connection.
func (c *Client) WritePump() {
	defer func() {
		c.conn.Close()
	}()
	for message := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}
