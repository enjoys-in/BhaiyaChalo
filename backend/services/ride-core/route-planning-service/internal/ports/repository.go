package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/route-planning-service/internal/model"
)

type RouteRepository interface {
	Save(ctx context.Context, route *model.Route) error
	FindByBookingID(ctx context.Context, bookingID string) (*model.Route, error)
}
