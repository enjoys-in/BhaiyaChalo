package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/incentive-service/internal/dto"
)

type IncentiveService interface {
	Create(ctx context.Context, req dto.CreateIncentiveRequest) (*dto.IncentiveResponse, error)
	GetByID(ctx context.Context, id string) (*dto.IncentiveResponse, error)
	Update(ctx context.Context, req dto.UpdateIncentiveRequest) (*dto.IncentiveResponse, error)
	Delete(ctx context.Context, id string) error
}
