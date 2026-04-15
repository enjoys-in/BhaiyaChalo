package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/support/support-ticket-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/support/support-ticket-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishTicketCreated(ctx context.Context, entity *model.Ticket) error {
	// TODO: publish to Kafka
	return nil
}

func (p *kafkaPublisher) PublishTicketResolved(ctx context.Context, entity *model.Ticket) error {
	// TODO: publish to Kafka
	return nil
}
