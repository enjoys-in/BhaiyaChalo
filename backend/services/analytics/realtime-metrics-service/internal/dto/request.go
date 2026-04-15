package dto

type CreateMetricRequest struct {
}

type UpdateMetricRequest struct {
	ID string `json:"id" validate:"required"`
}
