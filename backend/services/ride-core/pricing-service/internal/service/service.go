package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/pricing-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/pricing-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/pricing-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/pricing-service/internal/ports"
)

type pricingService struct {
	repo      ports.PricingRepository
	publisher ports.EventPublisher
}

func NewPricingService(repo ports.PricingRepository, publisher ports.EventPublisher) ports.PricingService {
	return &pricingService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *pricingService) Estimate(ctx context.Context, req dto.EstimatePriceRequest) (*dto.PriceEstimateResponse, error) {
	rule, err := s.repo.GetRule(ctx, req.CityID, req.VehicleType)
	if err != nil {
		return nil, fmt.Errorf("get pricing rule: %w", err)
	}
	if rule == nil {
		return nil, fmt.Errorf("no active pricing rule for city %s, vehicle %s", req.CityID, req.VehicleType)
	}

	surgeMultiplier := req.SurgeMultiplier
	if surgeMultiplier < 1.0 {
		surgeMultiplier = 1.0
	}

	baseFare := (req.DistanceKM * rule.BaseFarePerKM) + (req.DurationMin * rule.BaseFarePerMin) + rule.BookingFee
	estimatedFare := baseFare * surgeMultiplier

	if estimatedFare < rule.MinFare {
		estimatedFare = rule.MinFare
	}
	if rule.MaxFare > 0 && estimatedFare > rule.MaxFare {
		estimatedFare = rule.MaxFare
	}

	now := time.Now().UTC()
	estimate := &model.PriceEstimate{
		ID:              uuid.NewString(),
		CityID:          req.CityID,
		VehicleType:     req.VehicleType,
		DistanceKM:      req.DistanceKM,
		DurationMin:     req.DurationMin,
		BaseFare:        baseFare,
		SurgeMultiplier: surgeMultiplier,
		EstimatedFare:   estimatedFare,
		Currency:        constants.DefaultCurrency,
		CreatedAt:       now,
	}

	if err := s.repo.SaveEstimate(ctx, estimate); err != nil {
		return nil, fmt.Errorf("save estimate: %w", err)
	}

	_ = s.publisher.PublishPriceEstimated(ctx, estimate)

	return toEstimateResponse(estimate), nil
}

func (s *pricingService) GetRules(ctx context.Context, cityID string) ([]dto.PricingRuleResponse, error) {
	rules, err := s.repo.ListRules(ctx, cityID)
	if err != nil {
		return nil, fmt.Errorf("list rules: %w", err)
	}
	return toRuleResponseList(rules), nil
}

func (s *pricingService) CreateRule(ctx context.Context, req dto.CreatePricingRuleRequest) (*dto.PricingRuleResponse, error) {
	now := time.Now().UTC()
	rule := &model.PricingRule{
		ID:             uuid.NewString(),
		CityID:         req.CityID,
		VehicleType:    req.VehicleType,
		BaseFarePerKM:  req.BaseFarePerKM,
		BaseFarePerMin: req.BaseFarePerMin,
		MinFare:        req.MinFare,
		MaxFare:        req.MaxFare,
		BookingFee:     req.BookingFee,
		Active:         true,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := s.repo.CreateRule(ctx, rule); err != nil {
		return nil, fmt.Errorf("create rule: %w", err)
	}

	return toRuleResponse(rule), nil
}

func (s *pricingService) UpdateRule(ctx context.Context, id string, req dto.UpdatePricingRuleRequest) (*dto.PricingRuleResponse, error) {
	rules, err := s.repo.ListRules(ctx, "")
	if err != nil {
		return nil, fmt.Errorf("list rules: %w", err)
	}

	var rule *model.PricingRule
	for _, r := range rules {
		if r.ID == id {
			rule = r
			break
		}
	}
	if rule == nil {
		return nil, fmt.Errorf("pricing rule not found")
	}

	applyRuleUpdates(rule, req)
	rule.UpdatedAt = time.Now().UTC()

	if err := s.repo.UpdateRule(ctx, rule); err != nil {
		return nil, fmt.Errorf("update rule: %w", err)
	}

	_ = s.publisher.PublishRuleUpdated(ctx, rule)

	return toRuleResponse(rule), nil
}

func applyRuleUpdates(rule *model.PricingRule, req dto.UpdatePricingRuleRequest) {
	if req.BaseFarePerKM != nil {
		rule.BaseFarePerKM = *req.BaseFarePerKM
	}
	if req.BaseFarePerMin != nil {
		rule.BaseFarePerMin = *req.BaseFarePerMin
	}
	if req.MinFare != nil {
		rule.MinFare = *req.MinFare
	}
	if req.MaxFare != nil {
		rule.MaxFare = *req.MaxFare
	}
	if req.BookingFee != nil {
		rule.BookingFee = *req.BookingFee
	}
	if req.Active != nil {
		rule.Active = *req.Active
	}
}

func toEstimateResponse(e *model.PriceEstimate) *dto.PriceEstimateResponse {
	return &dto.PriceEstimateResponse{
		ID:              e.ID,
		CityID:          e.CityID,
		VehicleType:     e.VehicleType,
		DistanceKM:      e.DistanceKM,
		DurationMin:     e.DurationMin,
		BaseFare:        e.BaseFare,
		SurgeMultiplier: e.SurgeMultiplier,
		EstimatedFare:   e.EstimatedFare,
		Currency:        e.Currency,
		CreatedAt:       e.CreatedAt,
	}
}

func toRuleResponse(r *model.PricingRule) *dto.PricingRuleResponse {
	return &dto.PricingRuleResponse{
		ID:             r.ID,
		CityID:         r.CityID,
		VehicleType:    r.VehicleType,
		BaseFarePerKM:  r.BaseFarePerKM,
		BaseFarePerMin: r.BaseFarePerMin,
		MinFare:        r.MinFare,
		MaxFare:        r.MaxFare,
		BookingFee:     r.BookingFee,
		Active:         r.Active,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
	}
}

func toRuleResponseList(rules []*model.PricingRule) []dto.PricingRuleResponse {
	resp := make([]dto.PricingRuleResponse, len(rules))
	for i, r := range rules {
		resp[i] = *toRuleResponse(r)
	}
	return resp
}
