package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/user-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/user-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/user-service/internal/ports"
)

type eventEnvelope struct {
	EventType string    `json:"event_type"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}

type KafkaWriter interface {
	WriteMessages(ctx context.Context, topic string, key []byte, value []byte) error
}

type kafkaPublisher struct {
	writer KafkaWriter
}

func NewKafkaPublisher(writer KafkaWriter) ports.EventPublisher {
	return &kafkaPublisher{writer: writer}
}

func (p *kafkaPublisher) PublishUserCreated(ctx context.Context, user *model.User) error {
	return p.publish(ctx, constants.EventUserCreated, user)
}

func (p *kafkaPublisher) PublishUserUpdated(ctx context.Context, user *model.User) error {
	return p.publish(ctx, constants.EventUserUpdated, user)
}

func (p *kafkaPublisher) PublishUserDeleted(ctx context.Context, user *model.User) error {
	return p.publish(ctx, constants.EventUserDeleted, user)
}

func (p *kafkaPublisher) publish(ctx context.Context, eventType string, user *model.User) error {
	env := eventEnvelope{
		EventType: eventType,
		Timestamp: time.Now().UTC(),
		Payload:   user,
	}
	val, err := json.Marshal(env)
	if err != nil {
		return err
	}
	key := []byte(user.ID.String())
	return p.writer.WriteMessages(ctx, constants.KafkaTopicUserEvents, key, val)
}
