package broker

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Topic     string    `json:"topic"`
	Sender    string    `json:"sender"`
	Payload   string    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

type Subscriber struct {
	ID uuid.UUID
}
