package dto

type CreateReviewRequest struct {
}

type UpdateReviewRequest struct {
	ID string `json:"id" validate:"required"`
}
