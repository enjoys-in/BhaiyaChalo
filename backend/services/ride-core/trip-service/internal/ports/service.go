package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/trip-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/trip-service/internal/model"
)

type TripService interface {
	Create(ctx context.Context, req dto.CreateTripRequest) (*model.Trip, error)
	Get(ctx context.Context, tripID string) (*model.Trip, error)
	UpdateStatus(ctx context.Context, req dto.UpdateTripStatusRequest) error
	ListByUser(ctx context.Context, userID string, limit, offset int) ([]model.Trip, error)
	ListByDriver(ctx context.Context, driverID string, limit, offset int) ([]model.Trip, error)
}
