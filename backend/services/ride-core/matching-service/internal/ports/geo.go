package ports

import "context"

type DriverLocation struct {
	DriverID   string
	Lat        float64
	Lng        float64
	DistanceKM float64
}

type GeoIndex interface {
	AddDriver(ctx context.Context, driverID string, lat, lng float64) error
	RemoveDriver(ctx context.Context, driverID string) error
	FindNearby(ctx context.Context, lat, lng, radiusKM float64) ([]DriverLocation, error)
}
