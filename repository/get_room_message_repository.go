package repository

import (
	"context"
	"trocup-message/config"
	"trocup-message/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetMessagesByRoomID(roomID string) ([]models.Message, error) {
	var messages []models.Message

	cursor, err := config.MessageCollection.Find(context.Background(), bson.M{"roomID": roomID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
