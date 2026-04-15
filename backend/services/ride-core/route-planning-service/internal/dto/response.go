package dto

import "time"

type WaypointResponse struct {
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
	Order int     `json:"order"`
	Label string  `json:"label"`
}

type RouteResponse struct {
	ID          string             `json:"id"`
	BookingID   string             `json:"booking_id"`
	Waypoints   []WaypointResponse `json:"waypoints"`
	DistanceKM  float64            `json:"distance_km"`
	DurationMin float64            `json:"duration_min"`
	Polyline    string             `json:"polyline"`
	CreatedAt   time.Time          `json:"created_at"`
}
