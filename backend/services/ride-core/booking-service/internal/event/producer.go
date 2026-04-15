package event

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/booking-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/booking-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/booking-service/internal/ports"
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

func (p *kafkaPublisher) PublishBookingCreated(ctx context.Context, booking *model.Booking) error {
	return p.publish(ctx, constants.EventBookingCreated, booking)
}

func (p *kafkaPublisher) PublishBookingConfirmed(ctx context.Context, booking *model.Booking) error {
	return p.publish(ctx, constants.EventBookingConfirmed, booking)
}

func (p *kafkaPublisher) PublishBookingCancelled(ctx context.Context, bookingID string) error {
	return p.publish(ctx, constants.EventBookingCancelled, map[string]string{"booking_id": bookingID})
}

func (p *kafkaPublisher) PublishDriverAssigned(ctx context.Context, bookingID, driverID string) error {
	return p.publish(ctx, constants.EventDriverAssigned, map[string]string{
		"booking_id": bookingID,
		"driver_id":  driverID,
	})
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

	// TODO: integrate with actual Kafka producer (e.g. segmentio/kafka-go)
	_ = data
	_ = ctx

	return nil
}
