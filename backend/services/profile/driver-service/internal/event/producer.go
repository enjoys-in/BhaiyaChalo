package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/driver-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/driver-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/driver-service/internal/ports"
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

func (p *kafkaPublisher) PublishDriverCreated(ctx context.Context, driver *model.Driver) error {
	return p.publish(ctx, constants.EventDriverCreated, driver.ID, driver)
}

func (p *kafkaPublisher) PublishDriverUpdated(ctx context.Context, driver *model.Driver) error {
	return p.publish(ctx, constants.EventDriverUpdated, driver.ID, driver)
}

func (p *kafkaPublisher) PublishDriverStatusChanged(ctx context.Context, driver *model.Driver, oldStatus string) error {
	payload := map[string]interface{}{
		"driver_id":  driver.ID,
		"old_status": oldStatus,
		"new_status": driver.Status,
	}
	return p.publish(ctx, constants.EventDriverStatusChanged, driver.ID, payload)
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
