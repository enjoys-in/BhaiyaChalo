package dto

import "time"

type PriceEstimateResponse struct {
	ID              string    `json:"id"`
	CityID          string    `json:"city_id"`
	VehicleType     string    `json:"vehicle_type"`
	DistanceKM      float64   `json:"distance_km"`
	DurationMin     float64   `json:"duration_min"`
	BaseFare        float64   `json:"base_fare"`
	SurgeMultiplier float64   `json:"surge_multiplier"`
	EstimatedFare   float64   `json:"estimated_fare"`
	Currency        string    `json:"currency"`
	CreatedAt       time.Time `json:"created_at"`
}

type PricingRuleResponse struct {
	ID             string    `json:"id"`
	CityID         string    `json:"city_id"`
	VehicleType    string    `json:"vehicle_type"`
	BaseFarePerKM  float64   `json:"base_fare_per_km"`
	BaseFarePerMin float64   `json:"base_fare_per_min"`
	MinFare        float64   `json:"min_fare"`
	MaxFare        float64   `json:"max_fare"`
	BookingFee     float64   `json:"booking_fee"`
	Active         bool      `json:"active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
