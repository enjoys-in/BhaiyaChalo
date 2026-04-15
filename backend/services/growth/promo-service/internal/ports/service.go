package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/promo-service/internal/dto"
)

type PromoService interface {
	Create(ctx context.Context, req *dto.CreatePromoRequest) (*dto.PromoResponse, error)
	Apply(ctx context.Context, req *dto.ApplyPromoRequest) (*dto.ApplyPromoResponse, error)
	Validate(ctx context.Context, req *dto.ValidatePromoRequest) (*dto.ApplyPromoResponse, error)
	ListActive(ctx context.Context, cityID string) (*dto.PromoListResponse, error)
}
