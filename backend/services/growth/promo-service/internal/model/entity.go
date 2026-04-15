package model

import "time"

type PromoType string

const (
	PromoTypeFlat       PromoType = "flat"
	PromoTypePercentage PromoType = "percentage"
)

type PromoCode struct {
	ID            string    `json:"id" db:"id"`
	Code          string    `json:"code" db:"code"`
	CityID        string    `json:"city_id" db:"city_id"`
	Type          PromoType `json:"type" db:"type"`
	DiscountValue float64   `json:"discount_value" db:"discount_value"`
	MaxDiscount   float64   `json:"max_discount" db:"max_discount"`
	MinOrderValue float64   `json:"min_order_value" db:"min_order_value"`
	UsageLimit    int       `json:"usage_limit" db:"usage_limit"`
	UsedCount     int       `json:"used_count" db:"used_count"`
	ValidFrom     time.Time `json:"valid_from" db:"valid_from"`
	ValidUntil    time.Time `json:"valid_until" db:"valid_until"`
	Active        bool      `json:"active" db:"active"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

type PromoUsage struct {
	ID              string    `json:"id" db:"id"`
	PromoID         string    `json:"promo_id" db:"promo_id"`
	UserID          string    `json:"user_id" db:"user_id"`
	BookingID       string    `json:"booking_id" db:"booking_id"`
	DiscountApplied float64   `json:"discount_applied" db:"discount_applied"`
	UsedAt          time.Time `json:"used_at" db:"used_at"`
}
