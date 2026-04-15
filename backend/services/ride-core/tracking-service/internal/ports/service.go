package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/model"
)

type TrackingService interface {
	UpdateLocation(ctx context.Context, req *dto.UpdateLocationRequest) error
	GetLocation(ctx context.Context, driverID string) (*model.LocationUpdate, error)
	StartTracking(ctx context.Context, req *dto.StartTrackingRequest) (*model.TrackingSession, error)
	StopTracking(ctx context.Context, tripID string) error
}
