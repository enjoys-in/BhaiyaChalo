package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/risk-scoring-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/risk-scoring-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishRiskScored(ctx context.Context, entity *model.RiskScore) error {
	// TODO: publish to Kafka
	return nil
}

func (p *kafkaPublisher) PublishRiskFlagged(ctx context.Context, entity *model.RiskScore) error {
	// TODO: publish to Kafka
	return nil
}
