package event

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/earnings-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/earnings-service/internal/ports"
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
	return &kafkaPublisher{writer: writer, topic: topic}
}

func (p *kafkaPublisher) PublishEarningRecorded(ctx context.Context, driverID, tripID string, netEarning float64) error {
	payload := map[string]interface{}{
		"driver_id":   driverID,
		"trip_id":     tripID,
		"net_earning": netEarning,
	}
	return p.publish(ctx, constants.EventEarningRecorded, driverID, payload)
}

func (p *kafkaPublisher) publish(ctx context.Context, eventType, key string, payload interface{}) error {
	msg := message{
		EventType: eventType,
		Payload:   payload,
		Timestamp: time.Now().UTC(),
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal event: %w", err)
	}
	return p.writer.WriteMessages(ctx, p.topic, key, data)
}
