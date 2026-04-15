package ports

import "context"

type RouteResult struct {
	DistanceKM      float64
	DurationMinutes float64
}

type RoutingEngine interface {
	GetRoute(ctx context.Context, fromLat, fromLng, toLat, toLng float64) (*RouteResult, error)
}
