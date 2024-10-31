package services

import (
	"trocup-message/models"
	"trocup-message/repository"
)

func GetMessages() ([]models.Message, error) {
	return repository.GetMessages()
}
