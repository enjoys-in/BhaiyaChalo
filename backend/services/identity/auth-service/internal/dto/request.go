package dto

type LoginRequest struct {
	Phone string `json:"phone" validate:"required"`
	OTP   string `json:"otp" validate:"required"`
	Role  string `json:"role" validate:"required,oneof=user driver admin"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type VerifyRequest struct {
	AccessToken string `json:"access_token" validate:"required"`
}

type LogoutRequest struct {
	UserID string `json:"user_id"`
}

type SendOTPRequest struct {
	Phone string `json:"phone" validate:"required"`
}
