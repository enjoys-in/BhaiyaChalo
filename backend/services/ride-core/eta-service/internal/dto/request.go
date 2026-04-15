package dto

type CalculateETARequest struct {
	FromLat     float64 `json:"from_lat" validate:"required"`
	FromLng     float64 `json:"from_lng" validate:"required"`
	ToLat       float64 `json:"to_lat" validate:"required"`
	ToLng       float64 `json:"to_lng" validate:"required"`
	VehicleType string  `json:"vehicle_type" validate:"required"`
	CityID      string  `json:"city_id" validate:"required"`
}
