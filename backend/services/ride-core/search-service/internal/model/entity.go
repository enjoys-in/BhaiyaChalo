package model

import "time"

type SearchQuery struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	CityID    string    `json:"city_id" db:"city_id"`
	PickupLat float64   `json:"pickup_lat" db:"pickup_lat"`
	PickupLng float64   `json:"pickup_lng" db:"pickup_lng"`
	DropLat   float64   `json:"drop_lat" db:"drop_lat"`
	DropLng   float64   `json:"drop_lng" db:"drop_lng"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type SearchResult struct {
	QueryID          string  `json:"query_id" db:"query_id"`
	VehicleType      string  `json:"vehicle_type" db:"vehicle_type"`
	EstimatedFare    float64 `json:"estimated_fare" db:"estimated_fare"`
	ETAMinutes       int     `json:"eta_minutes" db:"eta_minutes"`
	AvailableDrivers int     `json:"available_drivers" db:"available_drivers"`
	SurgeMultiplier  float64 `json:"surge_multiplier" db:"surge_multiplier"`
}
