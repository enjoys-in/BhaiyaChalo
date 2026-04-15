package dto

type CreateRiskScoreRequest struct {
}

type UpdateRiskScoreRequest struct {
	ID string `json:"id" validate:"required"`
}
