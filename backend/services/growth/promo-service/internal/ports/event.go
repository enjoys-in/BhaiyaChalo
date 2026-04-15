package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/promo-service/internal/model"
)

type PromoEventPublisher interface {
	PublishPromoCreated(ctx context.Context, promo *model.PromoCode) error
	PublishPromoApplied(ctx context.Context, usage *model.PromoUsage) error
	PublishPromoExpired(ctx context.Context, promo *model.PromoCode) error
}
