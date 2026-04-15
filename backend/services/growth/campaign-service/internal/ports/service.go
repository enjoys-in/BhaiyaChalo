package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/campaign-service/internal/dto"
)

type CampaignService interface {
	Create(ctx context.Context, req *dto.CreateCampaignRequest) (*dto.CampaignResponse, error)
	Launch(ctx context.Context, campaignID string) (*dto.CampaignResponse, error)
	Pause(ctx context.Context, campaignID string) (*dto.CampaignResponse, error)
	GetStats(ctx context.Context, campaignID string) (*dto.CampaignStatsResponse, error)
	List(ctx context.Context, cityID string) (*dto.CampaignListResponse, error)
}
