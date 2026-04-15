package event

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/matching-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/matching-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/matching-service/internal/ports"
)

type kafkaPublisher struct {
	brokers string
}

func NewKafkaPublisher(brokers string) ports.EventPublisher {
	return &kafkaPublisher{brokers: brokers}
}

type eventEnvelope struct {
	Type      string      `json:"type"`
	Payload   interface{} `json:"payload"`
	Timestamp time.Time   `json:"timestamp"`
}

func (p *kafkaPublisher) PublishDriverMatched(ctx context.Context, result *model.MatchResult) error {
	return p.publish(ctx, constants.EventDriverMatched, result)
}

func (p *kafkaPublisher) PublishMatchFailed(ctx context.Context, bookingID string) error {
	return p.publish(ctx, constants.EventMatchFailed, map[string]string{"booking_id": bookingID})
}

func (p *kafkaPublisher) publish(ctx context.Context, eventType string, payload interface{}) error {
	envelope := eventEnvelope{
		Type:      eventType,
		Payload:   payload,
		Timestamp: time.Now().UTC(),
	}

	data, err := json.Marshal(envelope)
	if err != nil {
		return fmt.Errorf("marshal event: %w", err)
	}

	// TODO: integrate with actual Kafka producer
	_ = data
	_ = ctx

	return nil
}
