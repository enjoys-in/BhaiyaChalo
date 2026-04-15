package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/campaign-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/campaign-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/campaign-service/internal/ports"
)

type KafkaWriter interface {
	WriteMessage(ctx context.Context, topic string, key []byte, value []byte) error
}

type campaignEventPublisher struct {
	writer KafkaWriter
	logger *slog.Logger
}

func NewCampaignEventPublisher(writer KafkaWriter, logger *slog.Logger) ports.CampaignEventPublisher {
	return &campaignEventPublisher{writer: writer, logger: logger}
}

func (p *campaignEventPublisher) PublishCampaignLaunched(ctx context.Context, campaign *model.Campaign) error {
	return p.publish(ctx, constants.TopicCampaignLaunched, campaign.ID, campaign)
}

func (p *campaignEventPublisher) PublishCampaignCompleted(ctx context.Context, campaign *model.Campaign) error {
	return p.publish(ctx, constants.TopicCampaignCompleted, campaign.ID, campaign)
}

func (p *campaignEventPublisher) publish(ctx context.Context, topic, key string, payload interface{}) error {
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
