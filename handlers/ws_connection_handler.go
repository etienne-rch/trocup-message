// In handlers/connections.go
package handlers

import (
	"fmt"
	"net/http"
	"time"
	"trocup-message/models"
	"trocup-message/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Room struct {
	clients map[*websocket.Conn]bool
}

var rooms = make(map[string]*Room)
var broadcast = make(chan models.Message)

func HandleConnections(c *fiber.Ctx) error {
	roomID := c.Query("roomID")
	if roomID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("roomID is required")
	}

	// Convert Fiber context to standard http.Handler
	fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("WebSocket upgrade failed:", err)
			return
		}
		defer conn.Close()

		// Register the client in the specified room
		if rooms[roomID] == nil {
			rooms[roomID] = &Room{clients: make(map[*websocket.Conn]bool)}
		}
		rooms[roomID].clients[conn] = true

		// Handle incoming messages
		for {
			var msg models.Message
			err := conn.ReadJSON(&msg)
			if err != nil {
				fmt.Println("Error reading message:", err)
				delete(rooms[roomID].clients, conn)
				break
			}

			// Populate message fields
			msg.ID = primitive.NewObjectID()
			msg.SentAt = time.Now()
			msg.RoomID = roomID

			// Save message to database
			createdMessage, err := services.CreateMessage(&msg)
			if err != nil {
				fmt.Println("Error saving message to database:", err)
				continue
			}

			// Broadcast the message to the specific room
			broadcast <- *createdMessage
		}
	})(c.Context())

	return nil
}
