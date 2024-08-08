package tests

import (
	"testing"
	"time"
	"trocup-message/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMessage(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
	if err != nil {
		t.Fatalf("failed to create ObjectID: %v", err)
	}

	now := time.Now()

	message := models.Message{
		ID:       id,
		RoomID:   "room1",
		Sender:   id,
		Receiver: id,
		Message:  "Hello, World!",
		SentAt:   now,
	}

	if message.ID != id {
		t.Errorf("expected ID to be %v, got %v", id, message.ID)
	}
	if message.RoomID != "room1" {
		t.Errorf("expected RoomID to be 'room1', got %s", message.RoomID)
	}
	if message.Sender != id {
		t.Errorf("expected Sender to be %v, got %v", id, message.Sender)
	}
	if message.Receiver != id {
		t.Errorf("expected Receiver to be %v, got %v", id, message.Receiver)
	}
	if message.Message != "Hello, World!" {
		t.Errorf("expected Message to be 'Hello, World!', got %s", message.Message)
	}
	if !message.SentAt.Equal(now) {
		t.Errorf("expected SentAt to be %v, got %v", now, message.SentAt)
	}
}
