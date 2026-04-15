package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/iam-policy-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/iam-policy-service/internal/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.PolicyRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) CreateRole(ctx context.Context, role *model.Role) error {
	role.CreatedAt = time.Now().UTC()
	role.UpdatedAt = role.CreatedAt

	query := `INSERT INTO roles (id, name, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.ExecContext(ctx, query,
		role.ID, role.Name, role.Description, role.CreatedAt, role.UpdatedAt,
	)
	return err
}

func (r *postgresRepository) FindRoleByID(ctx context.Context, id string) (*model.Role, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM roles WHERE id = $1`

	var role model.Role
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *postgresRepository) ListRoles(ctx context.Context) ([]model.Role, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM roles ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []model.Role
	for rows.Next() {
		var role model.Role
		if err := rows.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, rows.Err()
}

func (r *postgresRepository) AddPermission(ctx context.Context, roleID, permissionID string) error {
	query := `INSERT INTO role_permissions (role_id, permission_id) VALUES ($1, $2)
		ON CONFLICT (role_id, permission_id) DO NOTHING`

	_, err := r.db.ExecContext(ctx, query, roleID, permissionID)
	return err
}

func (r *postgresRepository) FindPermissionsByRole(ctx context.Context, roleID string) ([]model.Permission, error) {
	query := `SELECT p.id, p.resource, p.action, p.description
		FROM permissions p
		INNER JOIN role_permissions rp ON rp.permission_id = p.id
		WHERE rp.role_id = $1`

	rows, err := r.db.QueryContext(ctx, query, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var perms []model.Permission
	for rows.Next() {
		var p model.Permission
		if err := rows.Scan(&p.ID, &p.Resource, &p.Action, &p.Description); err != nil {
			return nil, err
		}
		perms = append(perms, p)
	}
	return perms, rows.Err()
}

func (r *postgresRepository) CreatePolicy(ctx context.Context, policy *model.Policy) error {
	policy.CreatedAt = time.Now().UTC()

	query := `INSERT INTO policies (id, subject_type, subject_id, role_id, scope, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.ExecContext(ctx, query,
		policy.ID, policy.SubjectType, policy.SubjectID, policy.RoleID, policy.Scope, policy.CreatedAt,
	)
	return err
}

func (r *postgresRepository) CheckAccess(ctx context.Context, subjectType, subjectID, resource, action string) (bool, error) {
	query := `SELECT EXISTS (
		SELECT 1 FROM policies pol
		INNER JOIN role_permissions rp ON rp.role_id = pol.role_id
		INNER JOIN permissions p ON p.id = rp.permission_id
		WHERE pol.subject_type = $1
			AND pol.subject_id = $2
			AND p.resource = $3
			AND p.action = $4
	)`

	var allowed bool
	err := r.db.QueryRowContext(ctx, query, subjectType, subjectID, resource, action).Scan(&allowed)
	return allowed, err
}
