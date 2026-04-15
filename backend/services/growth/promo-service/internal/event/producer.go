package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/promo-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/promo-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/promo-service/internal/ports"
)

type KafkaWriter interface {
	WriteMessage(ctx context.Context, topic string, key []byte, value []byte) error
}

type promoEventPublisher struct {
	writer KafkaWriter
	logger *slog.Logger
}

func NewPromoEventPublisher(writer KafkaWriter, logger *slog.Logger) ports.PromoEventPublisher {
	return &promoEventPublisher{writer: writer, logger: logger}
}

func (p *promoEventPublisher) PublishPromoCreated(ctx context.Context, promo *model.PromoCode) error {
	return p.publish(ctx, constants.TopicPromoCreated, promo.ID, promo)
}

func (p *promoEventPublisher) PublishPromoApplied(ctx context.Context, usage *model.PromoUsage) error {
	return p.publish(ctx, constants.TopicPromoApplied, usage.PromoID, usage)
}

func (p *promoEventPublisher) PublishPromoExpired(ctx context.Context, promo *model.PromoCode) error {
	return p.publish(ctx, constants.TopicPromoExpired, promo.ID, promo)
}

func (p *promoEventPublisher) publish(ctx context.Context, topic, key string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshaling event: %w", err)
	}

	if err := p.writer.WriteMessage(ctx, topic, []byte(key), data); err != nil {
		p.logger.Error("failed to publish event", "topic", topic, "error", err)
		return fmt.Errorf("publishing to %s: %w", topic, err)
	}

	p.logger.Info("event published", "topic", topic, "key", key)
	return nil
}
