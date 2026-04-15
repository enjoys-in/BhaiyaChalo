package dto

import "time"

type TripResponse struct {
	ID             string     `json:"id"`
	BookingID      string     `json:"booking_id"`
	UserID         string     `json:"user_id"`
	DriverID       string     `json:"driver_id"`
	CityID         string     `json:"city_id"`
	VehicleType    string     `json:"vehicle_type"`
	Status         string     `json:"status"`
	PickupLat      float64    `json:"pickup_lat"`
	PickupLng      float64    `json:"pickup_lng"`
	DropLat        float64    `json:"drop_lat"`
	DropLng        float64    `json:"drop_lng"`
	ActualPickupAt *time.Time `json:"actual_pickup_at,omitempty"`
	ActualDropAt   *time.Time `json:"actual_drop_at,omitempty"`
	DistanceKM     float64    `json:"distance_km"`
	DurationMin    float64    `json:"duration_min"`
	FareAmount     float64    `json:"fare_amount"`
	PaymentMethod  string     `json:"payment_method"`
	Rating         float64    `json:"rating"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type TripTimelineResponse struct {
	TripID    string    `json:"trip_id"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
