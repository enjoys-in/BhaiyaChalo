package model

import "time"

type OTP struct {
	ID        string    `json:"id" db:"id"`
	Phone     string    `json:"phone" db:"phone"`
	Code      string    `json:"-" db:"code"`
	Purpose   Purpose   `json:"purpose" db:"purpose"`
	Verified  bool      `json:"verified" db:"verified"`
	Attempts  int       `json:"attempts" db:"attempts"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Purpose string

const (
	PurposeLogin         Purpose = "login"
	PurposePhoneVerify   Purpose = "phone_verify"
	PurposePasswordReset Purpose = "password_reset"
)
