package model

import "time"

type TripStatus string

const (
	TripStatusDriverEnroute TripStatus = "driver_enroute"
	TripStatusArrived       TripStatus = "arrived"
	TripStatusStarted       TripStatus = "trip_started"
	TripStatusCompleted     TripStatus = "completed"
	TripStatusCancelled     TripStatus = "cancelled"
)

type Trip struct {
	ID             string     `json:"id" bson:"_id"`
	BookingID      string     `json:"booking_id" bson:"booking_id"`
	UserID         string     `json:"user_id" bson:"user_id"`
	DriverID       string     `json:"driver_id" bson:"driver_id"`
	CityID         string     `json:"city_id" bson:"city_id"`
	VehicleType    string     `json:"vehicle_type" bson:"vehicle_type"`
	Status         TripStatus `json:"status" bson:"status"`
	PickupLat      float64    `json:"pickup_lat" bson:"pickup_lat"`
	PickupLng      float64    `json:"pickup_lng" bson:"pickup_lng"`
	DropLat        float64    `json:"drop_lat" bson:"drop_lat"`
	DropLng        float64    `json:"drop_lng" bson:"drop_lng"`
	ActualPickupAt *time.Time `json:"actual_pickup_at,omitempty" bson:"actual_pickup_at,omitempty"`
	ActualDropAt   *time.Time `json:"actual_drop_at,omitempty" bson:"actual_drop_at,omitempty"`
	DistanceKM     float64    `json:"distance_km" bson:"distance_km"`
	DurationMin    float64    `json:"duration_min" bson:"duration_min"`
	FareAmount     float64    `json:"fare_amount" bson:"fare_amount"`
	PaymentMethod  string     `json:"payment_method" bson:"payment_method"`
	Rating         float64    `json:"rating" bson:"rating"`
	CreatedAt      time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" bson:"updated_at"`
}

type TripTimeline struct {
	TripID    string    `json:"trip_id" bson:"trip_id"`
	Event     string    `json:"event" bson:"event"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
