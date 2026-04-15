package service

import (
	"context"
	"fmt"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/campaign-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/campaign-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/campaign-service/internal/ports"
	"github.com/google/uuid"
)

type campaignService struct {
	repo      ports.CampaignRepository
	publisher ports.CampaignEventPublisher
}

func NewCampaignService(repo ports.CampaignRepository, publisher ports.CampaignEventPublisher) ports.CampaignService {
	return &campaignService{repo: repo, publisher: publisher}
}

func (s *campaignService) Create(ctx context.Context, req *dto.CreateCampaignRequest) (*dto.CampaignResponse, error) {
	var scheduledAt *time.Time
	if req.ScheduledAt != "" {
		t, err := time.Parse(time.RFC3339, req.ScheduledAt)
		if err != nil {
			return nil, fmt.Errorf("invalid scheduled_at format: %w", err)
		}
		scheduledAt = &t
	}

	status := model.CampaignStatusDraft
	if scheduledAt != nil {
		status = model.CampaignStatusScheduled
	}

	campaign := &model.Campaign{
		ID:             uuid.New().String(),
		Name:           req.Name,
		Description:    req.Description,
		Type:           model.CampaignType(req.Type),
		TargetAudience: req.TargetAudience,
		CityID:         req.CityID,
		PromoCodeID:    req.PromoCodeID,
		Status:         status,
		ScheduledAt:    scheduledAt,
		CreatedAt:      time.Now(),
	}

	if err := s.repo.Create(ctx, campaign); err != nil {
		return nil, fmt.Errorf("creating campaign: %w", err)
	}

	return toCampaignResponse(campaign), nil
}

func (s *campaignService) Launch(ctx context.Context, campaignID string) (*dto.CampaignResponse, error) {
	campaign, err := s.repo.FindByID(ctx, campaignID)
	if err != nil {
		return nil, fmt.Errorf("finding campaign: %w", err)
	}
	if campaign == nil {
		return nil, fmt.Errorf("campaign not found")
	}

	if campaign.Status != model.CampaignStatusDraft && campaign.Status != model.CampaignStatusScheduled {
		return nil, fmt.Errorf("campaign cannot be launched from status: %s", campaign.Status)
	}

	now := time.Now()
	campaign.Status = model.CampaignStatusActive
	campaign.StartedAt = &now

	if err := s.repo.UpdateStatus(ctx, campaignID, model.CampaignStatusActive); err != nil {
		return nil, fmt.Errorf("updating campaign status: %w", err)
	}

	_ = s.publisher.PublishCampaignLaunched(ctx, campaign)

	return toCampaignResponse(campaign), nil
}

func (s *campaignService) Pause(ctx context.Context, campaignID string) (*dto.CampaignResponse, error) {
	campaign, err := s.repo.FindByID(ctx, campaignID)
	if err != nil {
		return nil, fmt.Errorf("finding campaign: %w", err)
	}
	if campaign == nil {
		return nil, fmt.Errorf("campaign not found")
	}

	if campaign.Status != model.CampaignStatusActive {
		return nil, fmt.Errorf("only active campaigns can be paused")
	}

	campaign.Status = model.CampaignStatusPaused

	if err := s.repo.UpdateStatus(ctx, campaignID, model.CampaignStatusPaused); err != nil {
		return nil, fmt.Errorf("updating campaign status: %w", err)
	}

	return toCampaignResponse(campaign), nil
}

func (s *campaignService) GetStats(ctx context.Context, campaignID string) (*dto.CampaignStatsResponse, error) {
	stats, err := s.repo.GetStats(ctx, campaignID)
	if err != nil {
		return nil, fmt.Errorf("getting stats: %w", err)
	}

	return &dto.CampaignStatsResponse{
		CampaignID: stats.CampaignID,
		Sent:       stats.Sent,
		Delivered:  stats.Delivered,
		Opened:     stats.Opened,
		Clicked:    stats.Clicked,
		Converted:  stats.Converted,
	}, nil
}

func (s *campaignService) List(ctx context.Context, cityID string) (*dto.CampaignListResponse, error) {
	campaigns, err := s.repo.List(ctx, cityID)
	if err != nil {
		return nil, fmt.Errorf("listing campaigns: %w", err)
	}

	responses := make([]dto.CampaignResponse, 0, len(campaigns))
	for i := range campaigns {
		responses = append(responses, *toCampaignResponse(&campaigns[i]))
	}

	return &dto.CampaignListResponse{Campaigns: responses, Total: len(responses)}, nil
}

func toCampaignResponse(c *model.Campaign) *dto.CampaignResponse {
	return &dto.CampaignResponse{
		ID:             c.ID,
		Name:           c.Name,
		Description:    c.Description,
		Type:           string(c.Type),
		TargetAudience: c.TargetAudience,
		CityID:         c.CityID,
		PromoCodeID:    c.PromoCodeID,
		Status:         string(c.Status),
		ScheduledAt:    dto.FormatTimePtr(c.ScheduledAt),
		StartedAt:      dto.FormatTimePtr(c.StartedAt),
		CompletedAt:    dto.FormatTimePtr(c.CompletedAt),
		CreatedAt:      dto.FormatTime(c.CreatedAt),
	}
}
