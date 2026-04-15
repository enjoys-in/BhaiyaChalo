package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/stop-planning-service/internal/model"
)

type StopRepository interface {
	AddStop(ctx context.Context, stop *model.Stop) error
	RemoveStop(ctx context.Context, tripID, stopID string) error
	ReorderStops(ctx context.Context, tripID string, stopIDs []string) error
	UpdateStopStatus(ctx context.Context, tripID, stopID string, status model.StopStatus) (*model.Stop, error)
	FindByTripID(ctx context.Context, tripID string) (*model.MultiStopTrip, error)
}
