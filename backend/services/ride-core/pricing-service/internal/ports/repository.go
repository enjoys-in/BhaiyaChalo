package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/pricing-service/internal/model"
)

type PricingRepository interface {
	SaveEstimate(ctx context.Context, estimate *model.PriceEstimate) error
	GetRule(ctx context.Context, cityID, vehicleType string) (*model.PricingRule, error)
	ListRules(ctx context.Context, cityID string) ([]*model.PricingRule, error)
	CreateRule(ctx context.Context, rule *model.PricingRule) error
	UpdateRule(ctx context.Context, rule *model.PricingRule) error
}
