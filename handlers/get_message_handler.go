package handlers

import (
	"trocup-message/services"

	"github.com/gofiber/fiber/v2"
)

func GetMessageByID(c *fiber.Ctx) error {
	id := c.Params("id")
	message, err := services.GetMessageByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Message not found"})
	}
	return c.JSON(message)
}
