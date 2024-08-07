package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
    RoomID   string             `bson:"roomID" json:"roomID"`
    Sender   primitive.ObjectID `bson:"sender" json:"sender"`
    Receiver primitive.ObjectID `bson:"receiver" json:"receiver"`
    Message  string             `bson:"message" json:"message"`
    SentAt   time.Time          `bson:"sentAt" json:"sentAt"`
}