package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/realtime-metrics-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/analytics/realtime-metrics-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishMetricsComputed(ctx context.Context, entity *model.Metric) error {
	// TODO: publish to Kafka
	return nil
}
