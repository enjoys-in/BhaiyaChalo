package dto

import "time"

type FareBreakdownResponse struct {
	ID              string    `json:"id"`
	BookingID       string    `json:"booking_id"`
	BasePrice       float64   `json:"base_price"`
	DistanceCharge  float64   `json:"distance_charge"`
	TimeCharge      float64   `json:"time_charge"`
	SurgeMultiplier float64   `json:"surge_multiplier"`
	SurgeAmount     float64   `json:"surge_amount"`
	TollCharges     float64   `json:"toll_charges"`
	TaxAmount       float64   `json:"tax_amount"`
	PromoDiscount   float64   `json:"promo_discount"`
	TotalFare       float64   `json:"total_fare"`
	Currency        string    `json:"currency"`
	CityID          string    `json:"city_id"`
	VehicleType     string    `json:"vehicle_type"`
	CreatedAt       time.Time `json:"created_at"`
}
