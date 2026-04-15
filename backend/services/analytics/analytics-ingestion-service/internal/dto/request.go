package dto

type CreateAnalyticsEventRequest struct {
}

type UpdateAnalyticsEventRequest struct {
	ID string `json:"id" validate:"required"`
}
