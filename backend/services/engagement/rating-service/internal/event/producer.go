package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/rating-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/engagement/rating-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishRatingSubmitted(ctx context.Context, entity *model.Rating) error {
	// TODO: publish to Kafka
	return nil
}

func (p *kafkaPublisher) PublishRatingUpdated(ctx context.Context, entity *model.Rating) error {
	// TODO: publish to Kafka
	return nil
}
