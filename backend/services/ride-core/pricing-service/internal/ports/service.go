package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/pricing-service/internal/dto"
)

type PricingService interface {
	Estimate(ctx context.Context, req dto.EstimatePriceRequest) (*dto.PriceEstimateResponse, error)
	GetRules(ctx context.Context, cityID string) ([]dto.PricingRuleResponse, error)
	CreateRule(ctx context.Context, req dto.CreatePricingRuleRequest) (*dto.PricingRuleResponse, error)
	UpdateRule(ctx context.Context, id string, req dto.UpdatePricingRuleRequest) (*dto.PricingRuleResponse, error)
}
