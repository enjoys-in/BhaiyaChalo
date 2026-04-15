package dto

import "time"

type StopResponse struct {
	ID         string     `json:"id"`
	TripID     string     `json:"trip_id"`
	Lat        float64    `json:"lat"`
	Lng        float64    `json:"lng"`
	Address    string     `json:"address"`
	StopOrder  int        `json:"stop_order"`
	Status     string     `json:"status"`
	ArrivedAt  *time.Time `json:"arrived_at,omitempty"`
	DepartedAt *time.Time `json:"departed_at,omitempty"`
}

type MultiStopResponse struct {
	TripID     string         `json:"trip_id"`
	Stops      []StopResponse `json:"stops"`
	TotalStops int            `json:"total_stops"`
}
