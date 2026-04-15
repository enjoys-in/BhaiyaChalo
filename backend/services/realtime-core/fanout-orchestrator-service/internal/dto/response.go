package dto

type FanoutResponse struct {
	MessageID   string `json:"message_id"`
	TargetCount int    `json:"target_count"`
	Status      string `json:"status"`
}

type DeliveryStatsResponse struct {
	MessageID     string  `json:"message_id"`
	TotalTargets  int64   `json:"total_targets"`
	Delivered     int64   `json:"delivered"`
	Pending       int64   `json:"pending"`
	Failed        int64   `json:"failed"`
	DeliveryRatio float64 `json:"delivery_ratio"`
}
