package dto

type CreateAuditEntryRequest struct {
}

type UpdateAuditEntryRequest struct {
	ID string `json:"id" validate:"required"`
}
