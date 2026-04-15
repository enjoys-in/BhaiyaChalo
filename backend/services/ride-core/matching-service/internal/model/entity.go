package model

import "time"

type MatchStatus string

const (
	MatchSearching MatchStatus = "searching"
	MatchMatched   MatchStatus = "matched"
	MatchFailed    MatchStatus = "failed"
)

type MatchRequest struct {
	ID          string      `json:"id" db:"id"`
	BookingID   string      `json:"booking_id" db:"booking_id"`
	CityID      string      `json:"city_id" db:"city_id"`
	PickupLat   float64     `json:"pickup_lat" db:"pickup_lat"`
	PickupLng   float64     `json:"pickup_lng" db:"pickup_lng"`
	VehicleType string      `json:"vehicle_type" db:"vehicle_type"`
	RadiusKM    float64     `json:"radius_km" db:"radius_km"`
	Status      MatchStatus `json:"status" db:"status"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
}

type MatchResult struct {
	RequestID  string  `json:"request_id"`
	DriverID   string  `json:"driver_id"`
	DistanceKM float64 `json:"distance_km"`
	ETASeconds int     `json:"eta_seconds"`
	Score      float64 `json:"score"`
}

type MatchConfig struct {
	CityID        string  `json:"city_id" db:"city_id"`
	MaxRadius     float64 `json:"max_radius" db:"max_radius"`
	MinScore      float64 `json:"min_score" db:"min_score"`
	MaxCandidates int     `json:"max_candidates" db:"max_candidates"`
}
