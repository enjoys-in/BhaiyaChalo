package dto

type CreateNotificationRequest struct {
}

type UpdateNotificationRequest struct {
	ID string `json:"id" validate:"required"`
}
