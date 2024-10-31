package repository

import (
	"context"
	"trocup-message/config"
	"trocup-message/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMessageByID(id string) (*models.Message, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var message models.Message
	err = config.MessageCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
