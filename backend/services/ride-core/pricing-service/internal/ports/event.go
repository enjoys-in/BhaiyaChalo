package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/pricing-service/internal/model"
)

type EventPublisher interface {
	PublishPriceEstimated(ctx context.Context, estimate *model.PriceEstimate) error
	PublishRuleUpdated(ctx context.Context, rule *model.PricingRule) error
}
