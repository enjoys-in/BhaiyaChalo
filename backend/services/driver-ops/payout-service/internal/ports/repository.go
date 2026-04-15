package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/payout-service/internal/model"
)

type PayoutRepository interface {
	Create(ctx context.Context, entity *model.Payout) error
	FindByID(ctx context.Context, id string) (*model.Payout, error)
	Update(ctx context.Context, entity *model.Payout) error
	Delete(ctx context.Context, id string) error
}
