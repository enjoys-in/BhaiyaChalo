package dto

type CreateTemplateRequest struct {
}

type UpdateTemplateRequest struct {
	ID string `json:"id" validate:"required"`
}
