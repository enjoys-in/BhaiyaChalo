package dto

import "time"

type LocationResponse struct {
	DriverID  string    `json:"driver_id"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	Heading   float64   `json:"heading"`
	Speed     float64   `json:"speed"`
	Accuracy  float64   `json:"accuracy"`
	Timestamp time.Time `json:"timestamp"`
}

type TrackingSessionResponse struct {
	ID        string     `json:"id"`
	TripID    string     `json:"trip_id"`
	DriverID  string     `json:"driver_id"`
	UserID    string     `json:"user_id"`
	Active    bool       `json:"active"`
	StartedAt time.Time  `json:"started_at"`
	EndedAt   *time.Time `json:"ended_at,omitempty"`
}

type NearbyDriver struct {
	DriverID string  `json:"driver_id"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	DistKM   float64 `json:"dist_km"`
}
