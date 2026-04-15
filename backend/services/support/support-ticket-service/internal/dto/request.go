package dto

type CreateTicketRequest struct {
}

type UpdateTicketRequest struct {
	ID string `json:"id" validate:"required"`
}
