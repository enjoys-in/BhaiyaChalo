package dto

type CreateEscalationRequest struct {
}

type UpdateEscalationRequest struct {
	ID string `json:"id" validate:"required"`
}
