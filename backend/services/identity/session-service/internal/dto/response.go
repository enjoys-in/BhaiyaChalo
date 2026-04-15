package dto

import "time"

type SessionResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Role      string    `json:"role"`
	DeviceID  string    `json:"device_id"`
	Active    bool      `json:"active"`
	ExpiresAt time.Time `json:"expires_at"`
}
