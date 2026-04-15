package dto

import "time"

type SessionResponse struct {
	ID             string     `json:"id"`
	UserID         string     `json:"user_id"`
	Role           string     `json:"role"`
	DeviceID       string     `json:"device_id"`
	ServerID       string     `json:"server_id"`
	ConnectedAt    time.Time  `json:"connected_at"`
	DisconnectedAt *time.Time `json:"disconnected_at,omitempty"`
	Active         bool       `json:"active"`
}

type ActiveSessionsResponse struct {
	Sessions []SessionResponse `json:"sessions"`
	Total    int64             `json:"total"`
}
