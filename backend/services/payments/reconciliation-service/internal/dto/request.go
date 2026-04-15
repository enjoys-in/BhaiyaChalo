package dto

type CreateReconciliationRequest struct {
}

type UpdateReconciliationRequest struct {
	ID string `json:"id" validate:"required"`
}
