package model

import "time"

type BookingStatus string

const (
	StatusPending        BookingStatus = "pending"
	StatusConfirmed      BookingStatus = "confirmed"
	StatusDriverAssigned BookingStatus = "driver_assigned"
	StatusInProgress     BookingStatus = "in_progress"
	StatusCompleted      BookingStatus = "completed"
	StatusCancelled      BookingStatus = "cancelled"
)

type Booking struct {
	ID             string        `json:"id" db:"id"`
	UserID         string        `json:"user_id" db:"user_id"`
	CityID         string        `json:"city_id" db:"city_id"`
	PickupLat      float64       `json:"pickup_lat" db:"pickup_lat"`
	PickupLng      float64       `json:"pickup_lng" db:"pickup_lng"`
	PickupAddress  string        `json:"pickup_address" db:"pickup_address"`
	DropLat        float64       `json:"drop_lat" db:"drop_lat"`
	DropLng        float64       `json:"drop_lng" db:"drop_lng"`
	DropAddress    string        `json:"drop_address" db:"drop_address"`
	VehicleType    string        `json:"vehicle_type" db:"vehicle_type"`
	EstimatedFare  float64       `json:"estimated_fare" db:"estimated_fare"`
	FinalFare      float64       `json:"final_fare" db:"final_fare"`
	PromoCode      string        `json:"promo_code" db:"promo_code"`
	DiscountAmount float64       `json:"discount_amount" db:"discount_amount"`
	Status         BookingStatus `json:"status" db:"status"`
	DriverID       string        `json:"driver_id" db:"driver_id"`
	PaymentMethod  string        `json:"payment_method" db:"payment_method"`
	CreatedAt      time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" db:"updated_at"`
}
