package repository

import (
	"context"
	"trocup-message/config"
	"trocup-message/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetMessages() ([]models.Message, error) {
	var messages []models.Message
	cursor, err := config.MessageCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var message models.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
