package handlers

import (
	"log"
	"trocup-message/models"
	"trocup-message/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateMessage(c *fiber.Ctx) error {
	var validate = validator.New()
	message := new(models.Message)

	// Parse le corps de la requête JSON dans le modèle Message
	if err := c.BodyParser(&message); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Récupérer l'ID utilisateur connecté à partir du contexte (défini par le middleware Clerk)
	clerkUserId := c.Locals("clerkUserId").(string)
	log.Printf("User connected: %s", clerkUserId)

	// Vérifier que l'ID du propriétaire dans le body est bien l'ID de l'utilisateur connecté
	if message.Sender != clerkUserId {
		log.Printf("User ID mismatch: message.Sender = %s, clerkUserId = %s", message.Sender, clerkUserId)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You do not have permission to create an message for this user"})
	}

	// Validation des données reçues via le validateur
	if err := validate.Struct(message); err != nil {
		log.Printf("Validation error: %v", err)
		// Retourner une erreur si la validation échoue
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Créer le message
	createdMessage, err := services.CreateMessage(message)
	if err != nil {
		log.Printf("Error creating message: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Retourner le message créé avec un statut 201 (Created)
	return c.Status(fiber.StatusCreated).JSON(createdMessage)

}
