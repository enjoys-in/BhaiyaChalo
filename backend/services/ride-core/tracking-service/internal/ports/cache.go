package ports

import "context"

type GeoLocation struct {
	DriverID string
	Lat      float64
	Lng      float64
	DistKM   float64
}

type LocationCache interface {
	Set(ctx context.Context, driverID string, lat, lng float64, ttlSeconds int) error
	Get(ctx context.Context, driverID string) (lat, lng float64, err error)
	GeoRadius(ctx context.Context, lat, lng, radiusKM float64) ([]GeoLocation, error)
}
