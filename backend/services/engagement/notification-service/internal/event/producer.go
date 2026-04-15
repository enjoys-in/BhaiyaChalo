package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/notification-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/engagement/notification-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishNotificationSent(ctx context.Context, entity *model.Notification) error {
	// TODO: publish to Kafka
	return nil
}

func (p *kafkaPublisher) PublishNotificationFailed(ctx context.Context, entity *model.Notification) error {
	// TODO: publish to Kafka
	return nil
}
