package handlers

import "fmt"

func HandleMessages() {
	for {
		// Receive message from the broadcast channel
		msg := <-broadcast

		// Retrieve the room
		roomID := msg.RoomID
		room, exists := rooms[roomID]
		if !exists {
			fmt.Printf("Room %s not found\n", roomID)
			continue
		}

		// Send message to all clients in the specified room
		for client := range room.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println("Error sending message:", err)
				client.Close()
				delete(room.clients, client)
			}
		}
	}
}
