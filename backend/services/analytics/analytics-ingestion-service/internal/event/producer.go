package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-ingestion-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-ingestion-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishEventIngested(ctx context.Context, entity *model.AnalyticsEvent) error {
	// TODO: publish to Kafka
	return nil
}
