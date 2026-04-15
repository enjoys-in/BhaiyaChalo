package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/ports"
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

func (p *kafkaPublisher) PublishDriverOnline(ctx context.Context, driverID, cityID, vehicleType string) error {
	payload := map[string]string{
		"driver_id":    driverID,
		"city_id":      cityID,
		"vehicle_type": vehicleType,
	}
	return p.publish(ctx, constants.EventDriverOnline, driverID, payload)
}

func (p *kafkaPublisher) PublishDriverOffline(ctx context.Context, driverID string) error {
	return p.publish(ctx, constants.EventDriverOffline, driverID, map[string]string{"driver_id": driverID})
}

func (p *kafkaPublisher) PublishDriverBusy(ctx context.Context, driverID string) error {
	return p.publish(ctx, constants.EventDriverBusy, driverID, map[string]string{"driver_id": driverID})
}

func (p *kafkaPublisher) PublishDriverFree(ctx context.Context, driverID string) error {
	return p.publish(ctx, constants.EventDriverFree, driverID, map[string]string{"driver_id": driverID})
}

func (p *kafkaPublisher) publish(ctx context.Context, eventType, key string, payload interface{}) error {
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
