package dto

type CreatePayoutRequest struct {
}

type UpdatePayoutRequest struct {
	ID string `json:"id" validate:"required"`
}
