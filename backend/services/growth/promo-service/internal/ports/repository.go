package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/promo-service/internal/model"
)

type PromoRepository interface {
	Create(ctx context.Context, promo *model.PromoCode) error
	FindByCode(ctx context.Context, code string) (*model.PromoCode, error)
	FindByID(ctx context.Context, id string) (*model.PromoCode, error)
	ListActive(ctx context.Context) ([]model.PromoCode, error)
	ListByCityID(ctx context.Context, cityID string) ([]model.PromoCode, error)
	IncrementUsage(ctx context.Context, promoID string) error
	RecordUsage(ctx context.Context, usage *model.PromoUsage) error
	HasUserUsed(ctx context.Context, promoID, userID string) (bool, error)
}
