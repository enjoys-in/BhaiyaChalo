package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/stop-planning-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/stop-planning-service/internal/model"
)

type StopService interface {
	AddStop(ctx context.Context, req dto.AddStopRequest) (*model.Stop, error)
	RemoveStop(ctx context.Context, req dto.RemoveStopRequest) error
	ReorderStops(ctx context.Context, req dto.ReorderStopsRequest) (*model.MultiStopTrip, error)
	UpdateStopStatus(ctx context.Context, req dto.UpdateStopStatusRequest) (*model.Stop, error)
	GetStopsByTrip(ctx context.Context, tripID string) (*model.MultiStopTrip, error)
}
