package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/delivery-ack-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/delivery-ack-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishDeliveryAck(ctx context.Context, entity *model.DeliveryAck) error {
	// TODO: publish to Kafka
	return nil
}

func (p *kafkaPublisher) PublishDeliveryTimeout(ctx context.Context, entity *model.DeliveryAck) error {
	// TODO: publish to Kafka
	return nil
}
