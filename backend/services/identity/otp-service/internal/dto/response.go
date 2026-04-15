package dto

type SendOTPResponse struct {
	ExpiresAt  int64 `json:"expires_at"`
	RetryAfter int   `json:"retry_after"`
}

type VerifyOTPResponse struct {
	Verified bool `json:"verified"`
}
