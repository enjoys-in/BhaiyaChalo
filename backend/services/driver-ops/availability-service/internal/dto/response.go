package dto

import "time"

type AvailabilityResponse struct {
	DriverID    string    `json:"driver_id"`
	CityID      string    `json:"city_id"`
	Online      bool      `json:"online"`
	OnTrip      bool      `json:"on_trip"`
	VehicleType string    `json:"vehicle_type"`
	Lat         float64   `json:"lat"`
	Lng         float64   `json:"lng"`
	LastSeenAt  time.Time `json:"last_seen_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type OnlineDriversResponse struct {
	CityID      string `json:"city_id"`
	VehicleType string `json:"vehicle_type"`
	Count       int    `json:"count"`
}
