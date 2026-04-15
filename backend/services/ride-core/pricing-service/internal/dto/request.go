package dto

type EstimatePriceRequest struct {
	CityID          string  `json:"city_id" validate:"required"`
	VehicleType     string  `json:"vehicle_type" validate:"required"`
	DistanceKM      float64 `json:"distance_km" validate:"required,gt=0"`
	DurationMin     float64 `json:"duration_min" validate:"required,gt=0"`
	SurgeMultiplier float64 `json:"surge_multiplier"`
}

type UpdatePricingRuleRequest struct {
	BaseFarePerKM  *float64 `json:"base_fare_per_km,omitempty"`
	BaseFarePerMin *float64 `json:"base_fare_per_min,omitempty"`
	MinFare        *float64 `json:"min_fare,omitempty"`
	MaxFare        *float64 `json:"max_fare,omitempty"`
	BookingFee     *float64 `json:"booking_fee,omitempty"`
	Active         *bool    `json:"active,omitempty"`
}

type CreatePricingRuleRequest struct {
	CityID         string  `json:"city_id" validate:"required"`
	VehicleType    string  `json:"vehicle_type" validate:"required"`
	BaseFarePerKM  float64 `json:"base_fare_per_km" validate:"required,gt=0"`
	BaseFarePerMin float64 `json:"base_fare_per_min" validate:"required,gt=0"`
	MinFare        float64 `json:"min_fare" validate:"required,gt=0"`
	MaxFare        float64 `json:"max_fare" validate:"required,gt=0"`
	BookingFee     float64 `json:"booking_fee"`
}
