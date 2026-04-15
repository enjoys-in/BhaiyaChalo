package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishEscalationCreated(ctx context.Context, entity *model.Escalation) error {
	// TODO: publish to Kafka
	return nil
}

func (p *kafkaPublisher) PublishEscalationResolved(ctx context.Context, entity *model.Escalation) error {
	// TODO: publish to Kafka
	return nil
}
