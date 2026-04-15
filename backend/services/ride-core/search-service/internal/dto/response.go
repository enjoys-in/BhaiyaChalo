package dto

type SearchResponse struct {
	QueryID string          `json:"query_id"`
	Results []VehicleOption `json:"results"`
}

type VehicleOption struct {
	VehicleType   string  `json:"vehicle_type"`
	Fare          float64 `json:"fare"`
	ETA           int     `json:"eta"`
	Surge         float64 `json:"surge"`
	DriversNearby int     `json:"drivers_nearby"`
}
