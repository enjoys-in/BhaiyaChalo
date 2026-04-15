package dto

import "time"

type BookingResponse struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	CityID         string    `json:"city_id"`
	PickupLat      float64   `json:"pickup_lat"`
	PickupLng      float64   `json:"pickup_lng"`
	PickupAddress  string    `json:"pickup_address"`
	DropLat        float64   `json:"drop_lat"`
	DropLng        float64   `json:"drop_lng"`
	DropAddress    string    `json:"drop_address"`
	VehicleType    string    `json:"vehicle_type"`
	EstimatedFare  float64   `json:"estimated_fare"`
	FinalFare      float64   `json:"final_fare"`
	PromoCode      string    `json:"promo_code,omitempty"`
	DiscountAmount float64   `json:"discount_amount,omitempty"`
	Status         string    `json:"status"`
	DriverID       string    `json:"driver_id,omitempty"`
	PaymentMethod  string    `json:"payment_method"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type BookingStatusResponse struct {
	BookingID string `json:"booking_id"`
	Status    string `json:"status"`
	DriverID  string `json:"driver_id,omitempty"`
}
