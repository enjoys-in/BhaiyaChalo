package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/iam-policy-service/internal/dto"
)

type PolicyService interface {
	CreateRole(ctx context.Context, req dto.CreateRoleRequest) (*dto.RoleResponse, error)
	GetRole(ctx context.Context, id string) (*dto.RoleResponse, error)
	ListRoles(ctx context.Context) ([]dto.RoleResponse, error)

	AssignPermission(ctx context.Context, req dto.AssignPermissionRequest) error
	GetPermissionsByRole(ctx context.Context, roleID string) ([]dto.PermissionResponse, error)

	CreatePolicy(ctx context.Context, req dto.CreatePolicyRequest) (*dto.PolicyResponse, error)
	CheckAccess(ctx context.Context, req dto.CheckAccessRequest) (*dto.AccessCheckResponse, error)
}
