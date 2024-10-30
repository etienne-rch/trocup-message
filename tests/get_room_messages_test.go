package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"trocup-message/config"
	"trocup-message/handlers"
	"trocup-message/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetRoomMessages(t *testing.T) {
	app := fiber.New()

	// Clean up the database before the test
	config.CleanUpTestDatabase("test_db")

	// Use the handler to retrieve messages by room ID
	app.Get("/messages/rooms/:id", handlers.GetMessagesByRoomID)

	// Prepare test messages for room ID "0"
	roomID := "0"
	messages := []models.Message{
		{
			RoomID:   roomID,
			Sender:   "clerk_user_id_12345",
			Receiver: "clerk_user_id_67890",
			Message:  "Hello clerk_user_id_67890!",
		},
		{
			RoomID:   roomID,
			Sender:   "clerk_user_id_67890",
			Receiver: "clerk_user_id_12345",
			Message:  "Hello clerk_user_id_12345!",
		},
		{
			RoomID:   roomID,
			Sender:   "clerk_user_id_12345",
			Receiver: "clerk_user_id_67890",
			Message:  "Bye!",
		},
	}

	// Insert messages into the database
	for _, message := range messages {
		_, err := config.MessageCollection.InsertOne(context.TODO(), message)
		assert.NoError(t, err, "Failed to insert message into the database")
	}

	// Create a request to get messages for the room
	req := httptest.NewRequest("GET", "/messages/rooms/"+roomID, nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err, "Failed to get response from server")
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Decode the JSON response to check the returned messages
	var response []models.Message
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err, "Failed to decode response body")

	// Validate the number of messages returned
	assert.Equal(t, len(messages), len(response), "Number of returned messages does not match expected count")

	// Validate the content of each message returned
	for i, message := range messages {
		assert.Equal(t, message.RoomID, response[i].RoomID)
		assert.Equal(t, message.Sender, response[i].Sender)
		assert.Equal(t, message.Receiver, response[i].Receiver)
		assert.Equal(t, message.Message, response[i].Message)
	}

	// Clean up the database after the test
	defer config.CleanUpTestDatabase("test_db")
}
