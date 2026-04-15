package dto

type CreateSessionRequest struct {
	UserID    string `json:"user_id"`
	Role      string `json:"role"`
	DeviceID  string `json:"device_id"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}

type GetSessionRequest struct {
	SessionID string `json:"session_id"`
}

type InvalidateRequest struct {
	SessionID string `json:"session_id"`
}

type InvalidateAllRequest struct {
	UserID string `json:"user_id"`
}
