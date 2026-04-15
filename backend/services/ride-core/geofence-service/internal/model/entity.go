package model

import "time"

type FenceType string

const (
	FenceTypeCityBoundary FenceType = "city_boundary"
	FenceTypeZone         FenceType = "zone"
	FenceTypeAirport      FenceType = "airport"
	FenceTypeRestricted   FenceType = "restricted"
	FenceTypeSurgeZone    FenceType = "surge_zone"
)

type Coordinate struct {
	Lat float64 `json:"lat" db:"lat"`
	Lng float64 `json:"lng" db:"lng"`
}

type Geofence struct {
	ID        string            `json:"id" db:"id"`
	CityID    string            `json:"city_id" db:"city_id"`
	Name      string            `json:"name" db:"name"`
	Type      FenceType         `json:"type" db:"type"`
	Polygon   []Coordinate      `json:"polygon" db:"polygon"`
	CenterLat float64           `json:"center_lat" db:"center_lat"`
	CenterLng float64           `json:"center_lng" db:"center_lng"`
	RadiusKM  float64           `json:"radius_km" db:"radius_km"`
	Active    bool              `json:"active" db:"active"`
	Metadata  map[string]string `json:"metadata" db:"metadata"`
	CreatedAt time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt time.Time         `json:"updated_at" db:"updated_at"`
}
