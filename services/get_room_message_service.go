package services

import (
	"trocup-message/models"
	"trocup-message/repository"
)

func GetMessagesByRoomID(roomID string) ([]models.Message, error) {
	return repository.GetMessagesByRoomID(roomID)
}
