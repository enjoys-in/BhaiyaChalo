package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/fare-service/internal/model"
)

type FareRepository interface {
	SaveCalculation(ctx context.Context, calc *model.FareCalculation) error
	FindByBookingID(ctx context.Context, bookingID string) (*model.FareCalculation, error)
	GetConfig(ctx context.Context, cityID, vehicleType string) (*model.FareConfig, error)
}
