package dto

type CreateQueryResultRequest struct {
}

type UpdateQueryResultRequest struct {
	ID string `json:"id" validate:"required"`
}
