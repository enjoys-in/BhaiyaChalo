package dto

import "time"

type CreatePromoRequest struct {
	Code          string  `json:"code" validate:"required,min=3,max=20"`
	CityID        string  `json:"city_id" validate:"required"`
	Type          string  `json:"type" validate:"required,oneof=flat percentage"`
	DiscountValue float64 `json:"discount_value" validate:"required,gt=0"`
	MaxDiscount   float64 `json:"max_discount" validate:"required,gt=0"`
	MinOrderValue float64 `json:"min_order_value" validate:"gte=0"`
	UsageLimit    int     `json:"usage_limit" validate:"required,gt=0"`
	ValidFrom     string  `json:"valid_from" validate:"required"`
	ValidUntil    string  `json:"valid_until" validate:"required"`
}

type ApplyPromoRequest struct {
	Code          string  `json:"code" validate:"required"`
	UserID        string  `json:"user_id" validate:"required"`
	BookingAmount float64 `json:"booking_amount" validate:"required,gt=0"`
	CityID        string  `json:"city_id" validate:"required"`
}

type ValidatePromoRequest struct {
	Code   string `json:"code" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
	CityID string `json:"city_id" validate:"required"`
}

func (r *CreatePromoRequest) ParseTimes() (validFrom, validUntil time.Time, err error) {
	validFrom, err = time.Parse(time.RFC3339, r.ValidFrom)
	if err != nil {
		return
	}
	validUntil, err = time.Parse(time.RFC3339, r.ValidUntil)
	return
}
