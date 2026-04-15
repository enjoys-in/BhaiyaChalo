package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/payout-service/internal/dto"
)

type PayoutService interface {
	Create(ctx context.Context, req dto.CreatePayoutRequest) (*dto.PayoutResponse, error)
	GetByID(ctx context.Context, id string) (*dto.PayoutResponse, error)
	Update(ctx context.Context, req dto.UpdatePayoutRequest) (*dto.PayoutResponse, error)
	Delete(ctx context.Context, id string) error
}
