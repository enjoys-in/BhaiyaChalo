package dto

type CreateLocationUpdateRequest struct {
}

type UpdateLocationUpdateRequest struct {
	ID string `json:"id" validate:"required"`
}
