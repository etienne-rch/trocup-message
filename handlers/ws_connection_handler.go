package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins
		return true
	},
}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

func HandleConnections(c *fiber.Ctx) error {
	// Convert Fiber's context to standard http.Handler
	fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Upgrade the connection to a WebSocket connection
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("WebSocket upgrade failed:", err)
			return
		}
		defer conn.Close()

		// Register the new client
		clients[conn] = true

		// Handle incoming messages
		for {
			var msg Message
			err := conn.ReadJSON(&msg)
			if err != nil {
				fmt.Println("Error reading message:", err)
				delete(clients, conn)
				break
			}
			// Broadcast the message
			broadcast <- msg
		}
	})(c.Context())

	return nil
}
