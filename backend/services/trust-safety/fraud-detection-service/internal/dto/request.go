package dto

type CreateFraudSignalRequest struct {
}

type UpdateFraudSignalRequest struct {
	ID string `json:"id" validate:"required"`
}
