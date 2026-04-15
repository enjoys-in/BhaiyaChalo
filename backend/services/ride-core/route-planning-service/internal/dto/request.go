package dto

type LatLng struct {
	Lat float64 `json:"lat" validate:"required,latitude"`
	Lng float64 `json:"lng" validate:"required,longitude"`
}

type PlanRouteRequest struct {
	BookingID   string   `json:"booking_id" validate:"required"`
	Origin      LatLng   `json:"origin" validate:"required"`
	Destination LatLng   `json:"destination" validate:"required"`
	Waypoints   []LatLng `json:"waypoints,omitempty"`
}
