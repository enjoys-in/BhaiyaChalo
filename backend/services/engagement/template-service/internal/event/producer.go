package event

import (
	"context"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/template-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/engagement/template-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}

func (p *kafkaPublisher) PublishTemplateCreated(ctx context.Context, entity *model.Template) error {
	// TODO: publish to Kafka
	return nil
}

func (p *kafkaPublisher) PublishTemplateUpdated(ctx context.Context, entity *model.Template) error {
	// TODO: publish to Kafka
	return nil
}
