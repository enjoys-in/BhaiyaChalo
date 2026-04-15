package dto

import "github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/model"

type SendOTPRequest struct {
	Phone   string        `json:"phone" validate:"required,e164"`
	Purpose model.Purpose `json:"purpose" validate:"required"`
}

type VerifyOTPRequest struct {
	Phone   string        `json:"phone" validate:"required,e164"`
	Code    string        `json:"code" validate:"required"`
	Purpose model.Purpose `json:"purpose" validate:"required"`
}
