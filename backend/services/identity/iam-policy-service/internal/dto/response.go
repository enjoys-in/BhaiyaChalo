package dto

import "time"

type RoleResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PermissionResponse struct {
	ID          string `json:"id"`
	Resource    string `json:"resource"`
	Action      string `json:"action"`
	Description string `json:"description"`
}

type PolicyResponse struct {
	ID          string    `json:"id"`
	SubjectType string    `json:"subject_type"`
	SubjectID   string    `json:"subject_id"`
	RoleID      string    `json:"role_id"`
	Scope       string    `json:"scope"`
	CreatedAt   time.Time `json:"created_at"`
}

type AccessCheckResponse struct {
	Allowed bool `json:"allowed"`
}
