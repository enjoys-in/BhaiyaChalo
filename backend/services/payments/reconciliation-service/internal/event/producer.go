package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishReconciliationStarted(ctx context.Context, entity *model.Reconciliation) error {
	// TODO: publish to Kafka
	return nil
}

func (p *kafkaPublisher) PublishReconciliationCompleted(ctx context.Context, entity *model.Reconciliation) error {
	// TODO: publish to Kafka
	return nil
}
