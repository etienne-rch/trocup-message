package services

import (
	"trocup-message/models"
	"trocup-message/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMessage(message *models.Message) error {
    return repository.CreateMessage(message)
}

func GetMessages() ([]models.Message, error) {
    return repository.GetMessages()
}

func GetMessageByID(id primitive.ObjectID) (*models.Message, error) {
    return repository.GetMessageByID(id)
}