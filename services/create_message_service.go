package services

import (
	"trocup-message/models"
	"trocup-message/repository"
)

func CreateMessage(message *models.Message) (*models.Message, error) {
	return repository.CreateMessage(message)
}
