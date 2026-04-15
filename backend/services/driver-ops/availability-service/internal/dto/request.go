package dto

type GoOnlineRequest struct {
	DriverID    string  `json:"driver_id" validate:"required"`
	CityID      string  `json:"city_id" validate:"required"`
	VehicleType string  `json:"vehicle_type" validate:"required"`
	Lat         float64 `json:"lat" validate:"required"`
	Lng         float64 `json:"lng" validate:"required"`
}

type GoOfflineRequest struct {
	DriverID string `json:"driver_id" validate:"required"`
}

type UpdateTripStatusRequest struct {
	DriverID string `json:"driver_id" validate:"required"`
	OnTrip   bool   `json:"on_trip"`
}
