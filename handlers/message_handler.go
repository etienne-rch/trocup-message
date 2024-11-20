package handlers

import (
	"net/http"
	"trocup-message/models"
	"trocup-message/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

// SetupRoutes configure les routes de l'API
func SetupRoutes(app *fiber.App) {
	app.Get("/messages", getMessages)      // Swagger documentation pour cette route ci-dessous
	app.Get("/messages/:id", getMessageByID) // Swagger documentation pour cette route ci-dessous
	app.Post("/messages", createMessage)  // Swagger documentation pour cette route ci-dessous
}

// @Summary Créer un nouveau message
// @Description Ajoute un nouveau message dans la base de données
// @Tags Messages
// @Accept json
// @Produce json
// @Param message body models.Message true "Détails du message"
// @Success 201 {object} models.Message
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /messages [post]
func createMessage(c *fiber.Ctx) error {
	message := new(models.Message)
	if err := c.BodyParser(message); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := services.CreateMessage(message)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(message)
}

// @Summary Obtenir tous les messages
// @Description Récupère la liste de tous les messages
// @Tags Messages
// @Produce json
// @Success 200 {array} models.Message
// @Failure 500 {object} map[string]interface{}
// @Router /messages [get]
func getMessages(c *fiber.Ctx) error {
	messages, err := services.GetMessages()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(messages)
}

// @Summary Obtenir un message par ID
// @Description Récupère un message à partir de son ID
// @Tags Messages
// @Produce json
// @Param id path string true "ID du message"
// @Success 200 {object} models.Message
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /messages/{id} [get]
func getMessageByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Invalid ID format",
		})
	}

	message, err := services.GetMessageByID(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(map[string]interface{}{
			"error": "Message not found",
		})
	}
	return c.JSON(message)
}
