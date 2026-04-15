package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/review-moderation-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/engagement/review-moderation-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishReviewSubmitted(ctx context.Context, entity *model.Review) error {
	// TODO: publish to Kafka
	return nil
}

func (p *kafkaPublisher) PublishReviewModerated(ctx context.Context, entity *model.Review) error {
	// TODO: publish to Kafka
	return nil
}
