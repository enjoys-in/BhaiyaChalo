package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/route-planning-service/internal/model"
)

type RoutingEngine interface {
	Route(ctx context.Context, waypoints []model.Waypoint) (*model.Route, error)
}
