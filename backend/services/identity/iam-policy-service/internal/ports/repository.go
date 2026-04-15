package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/iam-policy-service/internal/model"
)

type PolicyRepository interface {
	CreateRole(ctx context.Context, role *model.Role) error
	FindRoleByID(ctx context.Context, id string) (*model.Role, error)
	ListRoles(ctx context.Context) ([]model.Role, error)

	AddPermission(ctx context.Context, roleID, permissionID string) error
	FindPermissionsByRole(ctx context.Context, roleID string) ([]model.Permission, error)

	CreatePolicy(ctx context.Context, policy *model.Policy) error
	CheckAccess(ctx context.Context, subjectType, subjectID, resource, action string) (bool, error)
}
