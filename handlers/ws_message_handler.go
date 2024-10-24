package handlers

import "fmt"

func HandleMessages() {
	for {
		// Receive message from the broadcast channel
		msg := <-broadcast

		// Send message to all connected clients
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println("Error sending message:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
