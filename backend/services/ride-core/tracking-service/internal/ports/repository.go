package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/model"
)

type TrackingRepository interface {
	SaveLocation(ctx context.Context, loc *model.LocationUpdate) error
	GetLatestLocation(ctx context.Context, driverID string) (*model.LocationUpdate, error)
	StartSession(ctx context.Context, session *model.TrackingSession) error
	EndSession(ctx context.Context, tripID string) error
	FindActiveSession(ctx context.Context, tripID string) (*model.TrackingSession, error)
}
