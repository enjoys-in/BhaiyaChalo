package dto

import "time"

type ETAResponse struct {
	DistanceKM        float64   `json:"distance_km"`
	DurationMinutes   float64   `json:"duration_minutes"`
	TrafficMultiplier float64   `json:"traffic_multiplier"`
	CalculatedAt      time.Time `json:"calculated_at"`
	VehicleType       string    `json:"vehicle_type"`
	CityID            string    `json:"city_id"`
}
