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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetMessageByID(t *testing.T) {
	app := fiber.New()

	// Nettoyer la base de données avant le test
	config.CleanUpTestDatabase("test_db")

	// Utiliser le handler pour récupérer un article par ID
	app.Get("/messages/:id", handlers.GetMessageByID)

	message := models.Message{
		ID:       primitive.NewObjectID(),
		RoomID:   "999",
		Sender:   "clerk_user_id_12345",
		Receiver: "clerk_user_id_67890",
		Message:  "message",
	}

	_, err := config.MessageCollection.InsertOne(context.TODO(), message)
	assert.NoError(t, err)

	// Créer une requête GET pour récupérer le message par ID
	req := httptest.NewRequest("GET", "/messages/"+message.ID.Hex(), nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Décoder la réponse JSON pour vérifier le contenu de l'article
	var returnedMessage models.Message
	err = json.NewDecoder(resp.Body).Decode(&returnedMessage)
	assert.NoError(t, err)

	assert.Equal(t, message.RoomID, returnedMessage.RoomID)
	assert.Equal(t, message.Sender, returnedMessage.Sender)
	assert.Equal(t, message.Receiver, returnedMessage.Receiver)
	assert.Equal(t, message.Message, returnedMessage.Message)

	defer config.CleanUpTestDatabase("test_db")
}
