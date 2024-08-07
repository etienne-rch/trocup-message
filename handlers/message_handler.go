package handlers

import (
	"net/http"
	"trocup-message/models"
	"trocup-message/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/messages", getMessages)
    app.Get("/messages/:id", getMessageByID)
    app.Post("/messages", createMessage)
}

func createMessage(c *fiber.Ctx) error {
    message := new(models.Message)
    if err := c.BodyParser(message); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    err := services.CreateMessage(message)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(http.StatusCreated).JSON(message)
}

func getMessages(c *fiber.Ctx) error {
    messages, err := services.GetMessages()
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.JSON(messages)
}

func getMessageByID(c *fiber.Ctx) error {
    idParam := c.Params("id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ID format",
        })
    }

    message, err := services.GetMessageByID(id)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Message not found",
        })
    }
    return c.JSON(message)
}
