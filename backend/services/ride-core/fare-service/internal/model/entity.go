package model

import "time"

type FareCalculation struct {
	ID              string    `json:"id" db:"id"`
	BookingID       string    `json:"booking_id" db:"booking_id"`
	BasePrice       float64   `json:"base_price" db:"base_price"`
	DistanceCharge  float64   `json:"distance_charge" db:"distance_charge"`
	TimeCharge      float64   `json:"time_charge" db:"time_charge"`
	SurgeMultiplier float64   `json:"surge_multiplier" db:"surge_multiplier"`
	SurgeAmount     float64   `json:"surge_amount" db:"surge_amount"`
	TollCharges     float64   `json:"toll_charges" db:"toll_charges"`
	TaxAmount       float64   `json:"tax_amount" db:"tax_amount"`
	PromoDiscount   float64   `json:"promo_discount" db:"promo_discount"`
	TotalFare       float64   `json:"total_fare" db:"total_fare"`
	Currency        string    `json:"currency" db:"currency"`
	CityID          string    `json:"city_id" db:"city_id"`
	VehicleType     string    `json:"vehicle_type" db:"vehicle_type"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

type FareConfig struct {
	CityID          string  `json:"city_id" db:"city_id"`
	VehicleType     string  `json:"vehicle_type" db:"vehicle_type"`
	BasePricePerKM  float64 `json:"base_price_per_km" db:"base_price_per_km"`
	BasePricePerMin float64 `json:"base_price_per_min" db:"base_price_per_min"`
	MinFare         float64 `json:"min_fare" db:"min_fare"`
	CancellationFee float64 `json:"cancellation_fee" db:"cancellation_fee"`
}
