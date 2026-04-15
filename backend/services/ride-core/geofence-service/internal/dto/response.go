package dto

import "time"

type GeofenceResponse struct {
	ID        string            `json:"id"`
	CityID    string            `json:"city_id"`
	Name      string            `json:"name"`
	Type      string            `json:"type"`
	Polygon   []CoordinateDTO   `json:"polygon"`
	CenterLat float64           `json:"center_lat"`
	CenterLng float64           `json:"center_lng"`
	RadiusKM  float64           `json:"radius_km"`
	Active    bool              `json:"active"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type PointCheckResponse struct {
	InFence       bool               `json:"in_fence"`
	MatchedFences []GeofenceResponse `json:"matched_fences"`
}
