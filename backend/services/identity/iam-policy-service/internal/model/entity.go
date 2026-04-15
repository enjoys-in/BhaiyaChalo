package model

import "time"

type Role struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Permission struct {
	ID          string `json:"id" db:"id"`
	Resource    string `json:"resource" db:"resource"`
	Action      string `json:"action" db:"action"`
	Description string `json:"description" db:"description"`
}

type RolePermission struct {
	RoleID       string `json:"role_id" db:"role_id"`
	PermissionID string `json:"permission_id" db:"permission_id"`
}

type Policy struct {
	ID          string    `json:"id" db:"id"`
	SubjectType string    `json:"subject_type" db:"subject_type"`
	SubjectID   string    `json:"subject_id" db:"subject_id"`
	RoleID      string    `json:"role_id" db:"role_id"`
	Scope       string    `json:"scope" db:"scope"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
