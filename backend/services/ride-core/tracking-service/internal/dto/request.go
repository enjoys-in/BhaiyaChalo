package dto

type UpdateLocationRequest struct {
	DriverID string  `json:"driver_id" validate:"required"`
	Lat      float64 `json:"lat" validate:"required"`
	Lng      float64 `json:"lng" validate:"required"`
	Heading  float64 `json:"heading"`
	Speed    float64 `json:"speed"`
	Accuracy float64 `json:"accuracy"`
}

type GetLocationRequest struct {
	DriverID string `json:"driver_id" validate:"required"`
}

type StartTrackingRequest struct {
	TripID   string `json:"trip_id" validate:"required"`
	DriverID string `json:"driver_id" validate:"required"`
	UserID   string `json:"user_id" validate:"required"`
}

type StopTrackingRequest struct {
	TripID string `json:"trip_id" validate:"required"`
}
