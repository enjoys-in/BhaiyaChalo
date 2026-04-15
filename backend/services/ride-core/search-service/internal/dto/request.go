package dto

type SearchRequest struct {
	UserID    string  `json:"user_id" validate:"required"`
	PickupLat float64 `json:"pickup_lat" validate:"required"`
	PickupLng float64 `json:"pickup_lng" validate:"required"`
	DropLat   float64 `json:"drop_lat" validate:"required"`
	DropLng   float64 `json:"drop_lng" validate:"required"`
	CityID    string  `json:"city_id" validate:"required"`
}
