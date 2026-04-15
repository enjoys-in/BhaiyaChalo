package model

import "time"

type PriceEstimate struct {
	ID              string    `json:"id" db:"id"`
	CityID          string    `json:"city_id" db:"city_id"`
	VehicleType     string    `json:"vehicle_type" db:"vehicle_type"`
	DistanceKM      float64   `json:"distance_km" db:"distance_km"`
	DurationMin     float64   `json:"duration_min" db:"duration_min"`
	BaseFare        float64   `json:"base_fare" db:"base_fare"`
	SurgeMultiplier float64   `json:"surge_multiplier" db:"surge_multiplier"`
	EstimatedFare   float64   `json:"estimated_fare" db:"estimated_fare"`
	Currency        string    `json:"currency" db:"currency"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

type PricingRule struct {
	ID             string    `json:"id" db:"id"`
	CityID         string    `json:"city_id" db:"city_id"`
	VehicleType    string    `json:"vehicle_type" db:"vehicle_type"`
	BaseFarePerKM  float64   `json:"base_fare_per_km" db:"base_fare_per_km"`
	BaseFarePerMin float64   `json:"base_fare_per_min" db:"base_fare_per_min"`
	MinFare        float64   `json:"min_fare" db:"min_fare"`
	MaxFare        float64   `json:"max_fare" db:"max_fare"`
	BookingFee     float64   `json:"booking_fee" db:"booking_fee"`
	Active         bool      `json:"active" db:"active"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
