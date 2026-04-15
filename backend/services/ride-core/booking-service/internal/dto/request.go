package dto

type CreateBookingRequest struct {
	UserID        string  `json:"user_id" validate:"required"`
	CityID        string  `json:"city_id" validate:"required"`
	PickupLat     float64 `json:"pickup_lat" validate:"required"`
	PickupLng     float64 `json:"pickup_lng" validate:"required"`
	PickupAddress string  `json:"pickup_address" validate:"required"`
	DropLat       float64 `json:"drop_lat" validate:"required"`
	DropLng       float64 `json:"drop_lng" validate:"required"`
	DropAddress   string  `json:"drop_address" validate:"required"`
	VehicleType   string  `json:"vehicle_type" validate:"required"`
	PromoCode     string  `json:"promo_code,omitempty"`
	PaymentMethod string  `json:"payment_method" validate:"required"`
}

type CancelBookingRequest struct {
	BookingID string `json:"booking_id" validate:"required"`
	Reason    string `json:"reason,omitempty"`
}

type UpdateStatusRequest struct {
	BookingID string `json:"booking_id" validate:"required"`
	Status    string `json:"status" validate:"required"`
}
