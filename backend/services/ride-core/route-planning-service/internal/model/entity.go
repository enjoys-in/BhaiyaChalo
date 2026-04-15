package model

import "time"

type Waypoint struct {
	Lat   float64 `json:"lat" db:"lat"`
	Lng   float64 `json:"lng" db:"lng"`
	Order int     `json:"order" db:"order"`
	Label string  `json:"label" db:"label"`
}

type Route struct {
	ID          string     `json:"id" db:"id"`
	BookingID   string     `json:"booking_id" db:"booking_id"`
	Waypoints   []Waypoint `json:"waypoints" db:"waypoints"`
	DistanceKM  float64    `json:"distance_km" db:"distance_km"`
	DurationMin float64    `json:"duration_min" db:"duration_min"`
	Polyline    string     `json:"polyline" db:"polyline"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}
