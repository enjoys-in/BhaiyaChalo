package model

import "time"

type ETARequest struct {
	FromLat     float64 `json:"from_lat"`
	FromLng     float64 `json:"from_lng"`
	ToLat       float64 `json:"to_lat"`
	ToLng       float64 `json:"to_lng"`
	VehicleType string  `json:"vehicle_type"`
	CityID      string  `json:"city_id"`
}

type ETAResult struct {
	DistanceKM        float64   `json:"distance_km"`
	DurationMinutes   float64   `json:"duration_minutes"`
	TrafficMultiplier float64   `json:"traffic_multiplier"`
	CalculatedAt      time.Time `json:"calculated_at"`
}
