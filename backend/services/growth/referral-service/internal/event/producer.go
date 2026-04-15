package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/referral-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/referral-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/referral-service/internal/ports"
)

type KafkaWriter interface {
	WriteMessage(ctx context.Context, topic string, key []byte, value []byte) error
}

type referralEventPublisher struct {
	writer KafkaWriter
	logger *slog.Logger
}

func NewReferralEventPublisher(writer KafkaWriter, logger *slog.Logger) ports.ReferralEventPublisher {
	return &referralEventPublisher{writer: writer, logger: logger}
}

func (p *referralEventPublisher) PublishReferralApplied(ctx context.Context, referral *model.Referral) error {
	return p.publish(ctx, constants.TopicReferralApplied, referral.ID, referral)
}

func (p *referralEventPublisher) PublishReferralCompleted(ctx context.Context, referral *model.Referral) error {
	return p.publish(ctx, constants.TopicReferralCompleted, referral.ID, referral)
}

func (p *referralEventPublisher) PublishRewardCredited(ctx context.Context, reward *model.ReferralReward) error {
	return p.publish(ctx, constants.TopicRewardCredited, reward.ID, reward)
}

func (p *referralEventPublisher) publish(ctx context.Context, topic, key string, payload interface{}) error {
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
