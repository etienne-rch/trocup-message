package services

import (
	"trocup-message/models"
	"trocup-message/repository"
)

func GetMessageByID(id string) (*models.Message, error) {
	return repository.GetMessageByID(id)
}
