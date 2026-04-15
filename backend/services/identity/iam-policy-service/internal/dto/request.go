package dto

type CreateRoleRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type AssignPermissionRequest struct {
	RoleID       string `json:"role_id" validate:"required"`
	PermissionID string `json:"permission_id" validate:"required"`
}

type CreatePolicyRequest struct {
	SubjectType string `json:"subject_type" validate:"required"`
	SubjectID   string `json:"subject_id" validate:"required"`
	RoleID      string `json:"role_id" validate:"required"`
	Scope       string `json:"scope"`
}

type CheckAccessRequest struct {
	SubjectType string `json:"subject_type" validate:"required"`
	SubjectID   string `json:"subject_id" validate:"required"`
	Resource    string `json:"resource" validate:"required"`
	Action      string `json:"action" validate:"required"`
}
