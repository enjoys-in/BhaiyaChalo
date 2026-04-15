package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/trip-service/internal/model"
)

type TripRepository interface {
	Create(ctx context.Context, trip *model.Trip) error
	FindByID(ctx context.Context, id string) (*model.Trip, error)
	FindByUserID(ctx context.Context, userID string, limit, offset int) ([]model.Trip, error)
	FindByDriverID(ctx context.Context, driverID string, limit, offset int) ([]model.Trip, error)
	UpdateStatus(ctx context.Context, tripID string, status model.TripStatus) error
	AddTimelineEvent(ctx context.Context, event *model.TripTimeline) error
	GetTimeline(ctx context.Context, tripID string) ([]model.TripTimeline, error)
}
