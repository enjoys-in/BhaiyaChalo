package model

import "time"

type StopStatus string

const (
	StopStatusPending   StopStatus = "pending"
	StopStatusArrived   StopStatus = "arrived"
	StopStatusCompleted StopStatus = "completed"
	StopStatusSkipped   StopStatus = "skipped"
)

type Stop struct {
	ID         string     `json:"id" db:"id"`
	TripID     string     `json:"trip_id" db:"trip_id"`
	Lat        float64    `json:"lat" db:"lat"`
	Lng        float64    `json:"lng" db:"lng"`
	Address    string     `json:"address" db:"address"`
	StopOrder  int        `json:"stop_order" db:"stop_order"`
	Status     StopStatus `json:"status" db:"status"`
	ArrivedAt  *time.Time `json:"arrived_at,omitempty" db:"arrived_at"`
	DepartedAt *time.Time `json:"departed_at,omitempty" db:"departed_at"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
}

type MultiStopTrip struct {
	ID         string `json:"id" db:"id"`
	TripID     string `json:"trip_id" db:"trip_id"`
	Stops      []Stop `json:"stops"`
	TotalStops int    `json:"total_stops" db:"total_stops"`
}
