package tests

import (
	"bytes"
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

func TestCreateMessage(t *testing.T) {
	app := fiber.New()

	app.Post("/messages", func(c *fiber.Ctx) error {
		c.Locals("clerkUserId", "clerk_user_id_12345")
		c.Locals("clerkEmail", "john.doe@example.com")
		c.Locals("clerkName", "John")
		c.Locals("clerkSurname", "Doe")
		return handlers.CreateMessage(c)
	})

	message := models.Message{
		RoomID:   "999",
		Sender:   "clerk_user_id_12345",
		Receiver: "clerk_user_id_67890",
		Message:  "message",
	}

	jsonMessage, _ := json.Marshal(message)

	req := httptest.NewRequest("POST", "/messages", bytes.NewReader(jsonMessage))
	req.Header.Set("Content-type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	defer config.CleanUpTestDatabase("test_db")
}
