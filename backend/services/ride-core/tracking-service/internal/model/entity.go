package model

import "time"

type LocationUpdate struct {
	DriverID  string    `json:"driver_id"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	Heading   float64   `json:"heading"`
	Speed     float64   `json:"speed"`
	Accuracy  float64   `json:"accuracy"`
	Timestamp time.Time `json:"timestamp"`
}

type TrackingSession struct {
	ID        string     `json:"id" db:"id"`
	TripID    string     `json:"trip_id" db:"trip_id"`
	DriverID  string     `json:"driver_id" db:"driver_id"`
	UserID    string     `json:"user_id" db:"user_id"`
	Active    bool       `json:"active" db:"active"`
	StartedAt time.Time  `json:"started_at" db:"started_at"`
	EndedAt   *time.Time `json:"ended_at,omitempty" db:"ended_at"`
}
