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

func TestGetMessages(t *testing.T) {
	app := fiber.New()

	app.Get("/messages", handlers.GetMessages)

	messages := []models.Message{
		{
			RoomID:   "0",
			Sender:   "clerk_user_id_12345",
			Receiver: "clerk_user_id_67890",
			Message:  "Hello clerk_user_id_67890!",
		},
		{
			RoomID:   "0",
			Sender:   "clerk_user_id_67890",
			Receiver: "clerk_user_id_12345",
			Message:  "Hello clerk_user_id_12345!",
		},
	}

	for _, message := range messages {
		config.MessageCollection.InsertOne(context.TODO(), message)
	}

	req := httptest.NewRequest("GET", "/messages", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err, "Failed to get response from server")

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response []models.Message

	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err, "Failed to decode response body")

	assert.Equal(t, len(messages), len(response), "Number of returned messages does not match expected count")

	for i, message := range messages {
		assert.Equal(t, message.RoomID, response[i].RoomID)
		assert.Equal(t, message.Sender, response[i].Sender)
		assert.Equal(t, message.Receiver, response[i].Receiver)
		assert.Equal(t, message.Message, response[i].Message)
		// assert.Equal(t, message.SentAt, response.Messages[i].SentAt)
	}

	defer config.CleanUpTestDatabase("test_db")
}
