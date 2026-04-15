package dto

type AddStopRequest struct {
	TripID    string  `json:"trip_id" validate:"required"`
	Lat       float64 `json:"lat" validate:"required,latitude"`
	Lng       float64 `json:"lng" validate:"required,longitude"`
	Address   string  `json:"address" validate:"required"`
	StopOrder int     `json:"stop_order" validate:"gte=0"`
}

type RemoveStopRequest struct {
	TripID string `json:"trip_id" validate:"required"`
	StopID string `json:"stop_id" validate:"required"`
}

type ReorderStopsRequest struct {
	TripID  string   `json:"trip_id" validate:"required"`
	StopIDs []string `json:"stop_ids" validate:"required,min=1"`
}

type UpdateStopStatusRequest struct {
	TripID string `json:"trip_id" validate:"required"`
	StopID string `json:"stop_id" validate:"required"`
	Status string `json:"status" validate:"required,oneof=pending arrived completed skipped"`
}
