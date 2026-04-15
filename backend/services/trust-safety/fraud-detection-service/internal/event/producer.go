package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/fraud-detection-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/fraud-detection-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishFraudDetected(ctx context.Context, entity *model.FraudSignal) error {
	// TODO: publish to Kafka
	return nil
}

func (p *kafkaPublisher) PublishFraudCleared(ctx context.Context, entity *model.FraudSignal) error {
	// TODO: publish to Kafka
	return nil
}
