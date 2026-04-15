package dto

type CreateDispatchRequest struct {
	BookingID   string   `json:"booking_id" validate:"required"`
	CityID      string   `json:"city_id" validate:"required"`
	DriverIDs   []string `json:"driver_ids" validate:"required,min=1"`
	PickupLat   float64  `json:"pickup_lat" validate:"required"`
	PickupLng   float64  `json:"pickup_lng" validate:"required"`
	VehicleType string   `json:"vehicle_type" validate:"required"`
}

type DriverResponseRequest struct {
	OfferID  string `json:"offer_id" validate:"required"`
	Accepted bool   `json:"accepted"`
}
