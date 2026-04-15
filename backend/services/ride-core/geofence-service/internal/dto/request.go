package dto

type CreateGeofenceRequest struct {
	CityID    string            `json:"city_id" validate:"required"`
	Name      string            `json:"name" validate:"required"`
	Type      string            `json:"type" validate:"required,oneof=city_boundary zone airport restricted surge_zone"`
	Polygon   []CoordinateDTO   `json:"polygon" validate:"required,min=3"`
	CenterLat float64           `json:"center_lat" validate:"required"`
	CenterLng float64           `json:"center_lng" validate:"required"`
	RadiusKM  float64           `json:"radius_km"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

type UpdateGeofenceRequest struct {
	Name      string            `json:"name,omitempty"`
	Type      string            `json:"type,omitempty"`
	Polygon   []CoordinateDTO   `json:"polygon,omitempty"`
	CenterLat *float64          `json:"center_lat,omitempty"`
	CenterLng *float64          `json:"center_lng,omitempty"`
	RadiusKM  *float64          `json:"radius_km,omitempty"`
	Active    *bool             `json:"active,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

type PointInFenceRequest struct {
	Lat    float64 `json:"lat" validate:"required"`
	Lng    float64 `json:"lng" validate:"required"`
	CityID string  `json:"city_id"`
}

type CoordinateDTO struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
