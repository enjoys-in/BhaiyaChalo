package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/location-ingest-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/location-ingest-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishLocationUpdated(ctx context.Context, entity *model.LocationUpdate) error {
	// TODO: publish to Kafka
	return nil
}

func (p *kafkaPublisher) PublishLocationBatch(ctx context.Context, entity *model.LocationUpdate) error {
	// TODO: publish to Kafka
	return nil
}
