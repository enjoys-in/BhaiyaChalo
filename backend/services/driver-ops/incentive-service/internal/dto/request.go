package dto

type CreateIncentiveRequest struct {
}

type UpdateIncentiveRequest struct {
	ID string `json:"id" validate:"required"`
}
