package repository

import (
	"context"
	"trocup-message/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteMessage(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = config.MessageCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}
