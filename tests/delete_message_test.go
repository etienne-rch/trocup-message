package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"trocup-message/config"
	"trocup-message/handlers"
	"trocup-message/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDeleteMessage(t *testing.T) {
	app := fiber.New()

	// Nettoyer la base de données avant le test
	config.CleanUpTestDatabase("test_db")

	// Mock le middleware Clerk pour simuler l'authentification
	app.Use(func(c *fiber.Ctx) error {
		// Simuler un utilisateur authentifié en définissant un faux ID utilisateur dans le contexte
		c.Locals("clerkUserId", "clerk_user_id_67890") // ID simulé
		return c.Next()
	})

	app.Delete("/messages/:id", handlers.DeleteMessage)

	objectID, _ := primitive.ObjectIDFromHex("67221b4e7bb2933fda250d26") // utilisez un ID valide

	message := models.Message{
		ID:       objectID,
		RoomID:   "0",
		Sender:   "clerk_user_id_67890",
		Receiver: "clerk_user_id_12345",
		Message:  "Hello clerk_user_id_12345!",
	}
	result, err := config.MessageCollection.InsertOne(context.TODO(), message)
	assert.NoError(t, err)

	t.Log("Inserted message with ID:", result.InsertedID.(primitive.ObjectID).Hex())

	req := httptest.NewRequest("DELETE", "/messages/"+message.ID.Hex(), nil)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	var deletedMessage models.Message
	err = config.MessageCollection.FindOne(context.TODO(), bson.M{"_id": message.ID}).Decode(&deletedMessage)

	assert.Equal(t, mongo.ErrNoDocuments, err, "expected ErrNoDocuments, got %v", err)

	// Nettoyage après chaque test
	defer config.CleanUpTestDatabase("test_db")
}
