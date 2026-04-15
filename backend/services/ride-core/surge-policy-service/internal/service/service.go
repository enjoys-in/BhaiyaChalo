package service

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/surge-policy-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/surge-policy-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/surge-policy-service/internal/ports"
)

type surgeService struct {
	repo      ports.SurgeRepository
	publisher ports.EventPublisher
}

func NewSurgeService(repo ports.SurgeRepository, publisher ports.EventPublisher) ports.SurgeService {
	return &surgeService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *surgeService) Calculate(ctx context.Context, req dto.CalculateSurgeRequest) (*dto.SurgeResponse, error) {
	policy, err := s.repo.GetPolicy(ctx, req.CityID)
	if err != nil {
		return nil, fmt.Errorf("get surge policy: %w", err)
	}
	if policy == nil {
		return nil, fmt.Errorf("no active surge policy for city %s", req.CityID)
	}

	multiplier := calculateMultiplier(req.DemandCount, req.SupplyCount, policy)

	now := time.Now().UTC()
	zone := &model.SurgeZone{
		ID:                req.ZoneID,
		CityID:            req.CityID,
		GeofenceID:        req.ZoneID,
		CurrentMultiplier: multiplier,
		DemandCount:       req.DemandCount,
		SupplyCount:       req.SupplyCount,
		UpdatedAt:         now,
	}

	if err := s.repo.SaveZone(ctx, zone); err != nil {
		return nil, fmt.Errorf("save surge zone: %w", err)
	}

	history := &model.SurgeHistory{
		ID:           uuid.NewString(),
		ZoneID:       req.ZoneID,
		Multiplier:   multiplier,
		DemandCount:  req.DemandCount,
		SupplyCount:  req.SupplyCount,
		CalculatedAt: now,
	}
	if err := s.repo.SaveHistory(ctx, history); err != nil {
		return nil, fmt.Errorf("save surge history: %w", err)
	}

	_ = s.publisher.PublishSurgeUpdated(ctx, zone)

	return toSurgeResponse(zone), nil
}

func (s *surgeService) GetCurrentSurge(ctx context.Context, zoneID string) (*dto.SurgeResponse, error) {
	zone, err := s.repo.GetZone(ctx, zoneID)
	if err != nil {
		return nil, fmt.Errorf("get surge zone: %w", err)
	}
	if zone == nil {
		return nil, fmt.Errorf("surge zone not found")
	}
	return toSurgeResponse(zone), nil
}

func (s *surgeService) UpdatePolicy(ctx context.Context, cityID string, req dto.UpdatePolicyRequest) (*dto.SurgePolicyResponse, error) {
	policy, err := s.repo.GetPolicy(ctx, cityID)
	if err != nil {
		return nil, fmt.Errorf("get surge policy: %w", err)
	}
	if policy == nil {
		return nil, fmt.Errorf("surge policy not found for city %s", cityID)
	}

	applyPolicyUpdates(policy, req)

	if err := s.repo.UpdatePolicy(ctx, policy); err != nil {
		return nil, fmt.Errorf("update surge policy: %w", err)
	}

	_ = s.publisher.PublishSurgePolicyChanged(ctx, policy)

	return toPolicyResponse(policy), nil
}

// calculateMultiplier computes surge multiplier from demand/supply ratio.
// Uses stepped increments based on how far the ratio exceeds the minimum threshold.
func calculateMultiplier(demand, supply int, policy *model.SurgePolicy) float64 {
	if supply == 0 {
		return policy.MaxMultiplier
	}

	ratio := float64(demand) / float64(supply)
	if ratio <= policy.MinDemandSupplyRatio {
		return 1.0
	}

	excessRatio := ratio - policy.MinDemandSupplyRatio
	steps := math.Floor(excessRatio / policy.StepSize)
	multiplier := 1.0 + (steps * policy.StepSize)

	if multiplier > policy.MaxMultiplier {
		multiplier = policy.MaxMultiplier
	}
	return multiplier
}

func applyPolicyUpdates(policy *model.SurgePolicy, req dto.UpdatePolicyRequest) {
	if req.MinDemandSupplyRatio != nil {
		policy.MinDemandSupplyRatio = *req.MinDemandSupplyRatio
	}
	if req.MaxMultiplier != nil {
		policy.MaxMultiplier = *req.MaxMultiplier
	}
	if req.StepSize != nil {
		policy.StepSize = *req.StepSize
	}
	if req.CooldownMinutes != nil {
		policy.CooldownMinutes = *req.CooldownMinutes
	}
	if req.Active != nil {
		policy.Active = *req.Active
	}
}

func toSurgeResponse(z *model.SurgeZone) *dto.SurgeResponse {
	return &dto.SurgeResponse{
		ZoneID:            z.ID,
		CityID:            z.CityID,
		CurrentMultiplier: z.CurrentMultiplier,
		DemandCount:       z.DemandCount,
		SupplyCount:       z.SupplyCount,
		UpdatedAt:         z.UpdatedAt,
	}
}

func toPolicyResponse(p *model.SurgePolicy) *dto.SurgePolicyResponse {
	return &dto.SurgePolicyResponse{
		ID:                   p.ID,
		CityID:               p.CityID,
		MinDemandSupplyRatio: p.MinDemandSupplyRatio,
		MaxMultiplier:        p.MaxMultiplier,
		StepSize:             p.StepSize,
		CooldownMinutes:      p.CooldownMinutes,
		Active:               p.Active,
		CreatedAt:            p.CreatedAt,
	}
}
