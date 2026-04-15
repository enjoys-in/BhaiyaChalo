package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/route-planning-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/route-planning-service/internal/model"
)

type RouteService interface {
	PlanRoute(ctx context.Context, req dto.PlanRouteRequest) (*model.Route, error)
	GetRoute(ctx context.Context, bookingID string) (*model.Route, error)
}
