package dto

type CreateDeliveryAckRequest struct {
}

type UpdateDeliveryAckRequest struct {
	ID string `json:"id" validate:"required"`
}
