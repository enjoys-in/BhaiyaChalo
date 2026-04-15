package dto

type FindDriversRequest struct {
	BookingID   string  `json:"booking_id" validate:"required"`
	CityID      string  `json:"city_id" validate:"required"`
	PickupLat   float64 `json:"pickup_lat" validate:"required"`
	PickupLng   float64 `json:"pickup_lng" validate:"required"`
	VehicleType string  `json:"vehicle_type" validate:"required"`
	RadiusKM    float64 `json:"radius_km,omitempty"`
}
