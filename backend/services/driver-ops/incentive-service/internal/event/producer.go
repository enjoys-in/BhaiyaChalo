package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/incentive-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/incentive-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishIncentiveCreated(ctx context.Context, entity *model.Incentive) error {
	// TODO: publish to Kafka
	return nil
}

func (p *kafkaPublisher) PublishIncentiveCompleted(ctx context.Context, entity *model.Incentive) error {
	// TODO: publish to Kafka
	return nil
}
