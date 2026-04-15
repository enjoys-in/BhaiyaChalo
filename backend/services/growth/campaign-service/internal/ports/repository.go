package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/campaign-service/internal/model"
)

type CampaignRepository interface {
	Create(ctx context.Context, campaign *model.Campaign) error
	FindByID(ctx context.Context, id string) (*model.Campaign, error)
	List(ctx context.Context, cityID string) ([]model.Campaign, error)
	UpdateStatus(ctx context.Context, id string, status model.CampaignStatus) error
	SaveStats(ctx context.Context, stats *model.CampaignStats) error
	GetStats(ctx context.Context, campaignID string) (*model.CampaignStats, error)
}
