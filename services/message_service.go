package services

import (
	"trocup-message/models"
	"trocup-message/repositery"
)

// func CreateMessage(message *models.Message) error {
// 	return repository.CreateMessage(message)
// }

func GetMessages() ([]models.Message, error) {
	return repositery.GetMessages()
}

// func GetMessageByID(id primitive.ObjectID) (*models.Message, error) {
// 	return repository.GetMessageByID(id)
// }
