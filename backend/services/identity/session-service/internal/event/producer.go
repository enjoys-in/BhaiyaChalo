package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/session-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/session-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/session-service/internal/ports"
)

type message struct {
	EventType string      `json:"event_type"`
	Payload   interface{} `json:"payload"`
	Timestamp time.Time   `json:"timestamp"`
}

type KafkaWriter interface {
	WriteMessages(ctx context.Context, topic string, key string, value []byte) error
}

type kafkaPublisher struct {
	writer KafkaWriter
	topic  string
}

func NewKafkaPublisher(writer KafkaWriter, topic string) ports.EventPublisher {
	return &kafkaPublisher{
		writer: writer,
		topic:  topic,
	}
}

func (p *kafkaPublisher) PublishSessionCreated(ctx context.Context, session *model.Session) error {
	return p.publish(ctx, constants.EventSessionCreated, session.UserID, session)
}

func (p *kafkaPublisher) PublishSessionInvalidated(ctx context.Context, sessionID string, userID string) error {
	payload := map[string]string{
		"session_id": sessionID,
		"user_id":    userID,
	}
	return p.publish(ctx, constants.EventSessionInvalidated, userID, payload)
}

func (p *kafkaPublisher) publish(ctx context.Context, eventType string, key string, payload interface{}) error {
	msg := message{
		EventType: eventType,
		Payload:   payload,
		Timestamp: time.Now().UTC(),
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(ctx, p.topic, key, data)
}
