package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/audit-log-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/analytics/audit-log-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishAuditLogged(ctx context.Context, entity *model.AuditEntry) error {
	// TODO: publish to Kafka
	return nil
}
