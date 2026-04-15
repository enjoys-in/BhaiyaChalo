package dto

type CreateTripRequest struct {
	BookingID     string  `json:"booking_id" validate:"required"`
	UserID        string  `json:"user_id" validate:"required"`
	DriverID      string  `json:"driver_id" validate:"required"`
	CityID        string  `json:"city_id" validate:"required"`
	VehicleType   string  `json:"vehicle_type" validate:"required"`
	PickupLat     float64 `json:"pickup_lat" validate:"required,latitude"`
	PickupLng     float64 `json:"pickup_lng" validate:"required,longitude"`
	DropLat       float64 `json:"drop_lat" validate:"required,latitude"`
	DropLng       float64 `json:"drop_lng" validate:"required,longitude"`
	FareAmount    float64 `json:"fare_amount" validate:"gte=0"`
	PaymentMethod string  `json:"payment_method" validate:"required"`
}

type UpdateTripStatusRequest struct {
	TripID string `json:"trip_id" validate:"required"`
	Status string `json:"status" validate:"required,oneof=driver_enroute arrived trip_started completed cancelled"`
}

type RateTripRequest struct {
	TripID string  `json:"trip_id" validate:"required"`
	Rating float64 `json:"rating" validate:"required,gte=1,lte=5"`
}
