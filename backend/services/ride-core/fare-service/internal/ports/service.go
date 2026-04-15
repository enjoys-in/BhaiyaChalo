package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/fare-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/fare-service/internal/model"
)

type FareService interface {
	Calculate(ctx context.Context, req dto.CalculateFareRequest) (*model.FareCalculation, error)
	Recalculate(ctx context.Context, req dto.RecalculateFareRequest) (*model.FareCalculation, error)
	GetBreakdown(ctx context.Context, bookingID string) (*model.FareCalculation, error)
}
