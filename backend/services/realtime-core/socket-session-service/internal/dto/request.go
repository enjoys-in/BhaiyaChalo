package dto

type RegisterSessionRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=user driver admin"`
	DeviceID string `json:"device_id" validate:"required"`
	ServerID string `json:"server_id" validate:"required"`
}

type UnregisterSessionRequest struct {
	SessionID string `json:"session_id" validate:"required"`
	UserID    string `json:"user_id" validate:"required"`
}
