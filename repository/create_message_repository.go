package repository

import (
	"context"
	"trocup-message/config"
	"trocup-message/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMessage(article *models.Message) (*models.Message, error) {
	result, err := config.MessageCollection.InsertOne(context.Background(), article)
	if err != nil {
		return nil, err
	}
	article.ID = result.InsertedID.(primitive.ObjectID)
	return article, nil
}
