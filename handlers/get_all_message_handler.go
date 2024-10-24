package handlers

import (
	"net/http"
	"trocup-message/services"

	"github.com/gofiber/fiber/v2"
)

func GetMessages(c *fiber.Ctx) error {
	messages, err := services.GetMessages()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(messages)
}
