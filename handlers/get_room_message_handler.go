package handlers

import (
	"trocup-message/services"

	"github.com/gofiber/fiber/v2"
)

func GetMessagesByRoomID(c *fiber.Ctx) error {
	roomID := c.Params("id")
	messages, err := services.GetMessagesByRoomID(roomID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Messages not found"})
	}
	return c.JSON(messages)
}
