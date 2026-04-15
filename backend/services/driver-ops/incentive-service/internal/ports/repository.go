package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/incentive-service/internal/model"
)

type IncentiveRepository interface {
	Create(ctx context.Context, entity *model.Incentive) error
	FindByID(ctx context.Context, id string) (*model.Incentive, error)
	Update(ctx context.Context, entity *model.Incentive) error
	Delete(ctx context.Context, id string) error
}
