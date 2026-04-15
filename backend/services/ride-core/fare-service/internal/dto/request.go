package dto

type CalculateFareRequest struct {
	BookingID       string  `json:"booking_id" validate:"required"`
	DistanceKM      float64 `json:"distance_km" validate:"required,gt=0"`
	DurationMin     float64 `json:"duration_min" validate:"required,gt=0"`
	CityID          string  `json:"city_id" validate:"required"`
	VehicleType     string  `json:"vehicle_type" validate:"required"`
	SurgeMultiplier float64 `json:"surge_multiplier" validate:"gte=1"`
	PromoDiscount   float64 `json:"promo_discount" validate:"gte=0"`
	TollCharges     float64 `json:"toll_charges" validate:"gte=0"`
}

type RecalculateFareRequest struct {
	BookingID       string  `json:"booking_id" validate:"required"`
	DistanceKM      float64 `json:"distance_km" validate:"required,gt=0"`
	DurationMin     float64 `json:"duration_min" validate:"required,gt=0"`
	SurgeMultiplier float64 `json:"surge_multiplier" validate:"gte=1"`
	PromoDiscount   float64 `json:"promo_discount" validate:"gte=0"`
	TollCharges     float64 `json:"toll_charges" validate:"gte=0"`
}
