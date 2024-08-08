package repository

import (
	"context"
	"trocup-message/config"
	"trocup-message/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MessageCollection *mongo.Collection

func InitMessageRepository() {
    MessageCollection = config.Client.Database("message_dev").Collection("messages")
}

func CreateMessage(message *models.Message) error {
    _, err := MessageCollection.InsertOne(context.TODO(), message)
    return err
}

func GetMessages() ([]models.Message, error) {
    var messages []models.Message
    cursor, err := MessageCollection.Find(context.TODO(), bson.M{})
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

func GetMessageByID(id primitive.ObjectID) (*models.Message, error) {
    var article models.Message
    err := MessageCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&article)
    if err != nil {
        return nil, err
    }
    return &article, nil
}
