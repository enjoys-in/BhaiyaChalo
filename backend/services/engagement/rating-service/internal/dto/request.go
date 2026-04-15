package dto

type CreateRatingRequest struct {
}

type UpdateRatingRequest struct {
	ID string `json:"id" validate:"required"`
}
