package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/campaign-service/internal/model"
)

type CampaignEventPublisher interface {
	PublishCampaignLaunched(ctx context.Context, campaign *model.Campaign) error
	PublishCampaignCompleted(ctx context.Context, campaign *model.Campaign) error
}
