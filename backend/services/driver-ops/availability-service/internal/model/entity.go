package model

import "time"

type ActionType string

const (
	ActionWentOnline  ActionType = "went_online"
	ActionWentOffline ActionType = "went_offline"
	ActionTripStarted ActionType = "trip_started"
	ActionTripEnded   ActionType = "trip_ended"
)

type DriverAvailability struct {
	DriverID    string    `json:"driver_id" db:"driver_id"`
	CityID      string    `json:"city_id" db:"city_id"`
	Online      bool      `json:"online" db:"online"`
	OnTrip      bool      `json:"on_trip" db:"on_trip"`
	VehicleType string    `json:"vehicle_type" db:"vehicle_type"`
	Lat         float64   `json:"lat" db:"lat"`
	Lng         float64   `json:"lng" db:"lng"`
	LastSeenAt  time.Time `json:"last_seen_at" db:"last_seen_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type AvailabilityLog struct {
	ID        string     `json:"id" db:"id"`
	DriverID  string     `json:"driver_id" db:"driver_id"`
	Action    ActionType `json:"action" db:"action"`
	Timestamp time.Time  `json:"timestamp" db:"timestamp"`
}
