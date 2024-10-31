package handlers

import (
	"trocup-message/services"

	"github.com/gofiber/fiber/v2"
)

func DeleteMessage(c *fiber.Ctx) error {
	// Récupérer l'ID du message à partir des paramètres d'URL
	id := c.Params("id")

	// Récupérer l'ID utilisateur connecté à partir du contexte (défini par le middleware Clerk)
	clerkUserId := c.Locals("clerkUserId").(string)

	// Récupérer le message à partir de la base de données via le service GetMessageByID
	message, err := services.GetMessageByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Message not found"})
	}

	// Vérifier que l'utilisateur connecté est bien le propriétaire du message
	if message.Sender != clerkUserId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You do not have permission to delete this message"})
	}

	// Supprimer le message via le service
	err = services.DeleteMessage(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Retourner un statut 204 No Content si la suppression est réussie
	return c.SendStatus(fiber.StatusNoContent)
}
