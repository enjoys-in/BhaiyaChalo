package service

import (
	"context"
	"fmt"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/iam-policy-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/iam-policy-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/iam-policy-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/iam-policy-service/internal/ports"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type policyService struct {
	repo      ports.PolicyRepository
	publisher ports.EventPublisher
	cache     *redis.Client
}

func NewPolicyService(repo ports.PolicyRepository, publisher ports.EventPublisher, cache *redis.Client) ports.PolicyService {
	return &policyService{
		repo:      repo,
		publisher: publisher,
		cache:     cache,
	}
}

func (s *policyService) CreateRole(ctx context.Context, req dto.CreateRoleRequest) (*dto.RoleResponse, error) {
	role := &model.Role{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.repo.CreateRole(ctx, role); err != nil {
		return nil, fmt.Errorf("create role: %w", err)
	}

	_ = s.publisher.PublishRoleCreated(ctx, role.ID, role.Name)

	return toRoleResponse(role), nil
}

func (s *policyService) GetRole(ctx context.Context, id string) (*dto.RoleResponse, error) {
	role, err := s.repo.FindRoleByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find role: %w", err)
	}
	return toRoleResponse(role), nil
}

func (s *policyService) ListRoles(ctx context.Context) ([]dto.RoleResponse, error) {
	roles, err := s.repo.ListRoles(ctx)
	if err != nil {
		return nil, fmt.Errorf("list roles: %w", err)
	}

	resp := make([]dto.RoleResponse, len(roles))
	for i, r := range roles {
		resp[i] = *toRoleResponse(&r)
	}
	return resp, nil
}

func (s *policyService) AssignPermission(ctx context.Context, req dto.AssignPermissionRequest) error {
	if err := s.repo.AddPermission(ctx, req.RoleID, req.PermissionID); err != nil {
		return fmt.Errorf("assign permission: %w", err)
	}
	return nil
}

func (s *policyService) GetPermissionsByRole(ctx context.Context, roleID string) ([]dto.PermissionResponse, error) {
	perms, err := s.repo.FindPermissionsByRole(ctx, roleID)
	if err != nil {
		return nil, fmt.Errorf("find permissions: %w", err)
	}

	resp := make([]dto.PermissionResponse, len(perms))
	for i, p := range perms {
		resp[i] = dto.PermissionResponse{
			ID:          p.ID,
			Resource:    p.Resource,
			Action:      p.Action,
			Description: p.Description,
		}
	}
	return resp, nil
}

func (s *policyService) CreatePolicy(ctx context.Context, req dto.CreatePolicyRequest) (*dto.PolicyResponse, error) {
	policy := &model.Policy{
		ID:          uuid.New().String(),
		SubjectType: req.SubjectType,
		SubjectID:   req.SubjectID,
		RoleID:      req.RoleID,
		Scope:       req.Scope,
	}

	if err := s.repo.CreatePolicy(ctx, policy); err != nil {
		return nil, fmt.Errorf("create policy: %w", err)
	}

	s.invalidateAccessCache(ctx, req.SubjectType, req.SubjectID)
	_ = s.publisher.PublishPolicyChanged(ctx, policy.ID, policy.SubjectType, policy.SubjectID)

	return &dto.PolicyResponse{
		ID:          policy.ID,
		SubjectType: policy.SubjectType,
		SubjectID:   policy.SubjectID,
		RoleID:      policy.RoleID,
		Scope:       policy.Scope,
		CreatedAt:   policy.CreatedAt,
	}, nil
}

func (s *policyService) CheckAccess(ctx context.Context, req dto.CheckAccessRequest) (*dto.AccessCheckResponse, error) {
	cacheKey := fmt.Sprintf("%s%s:%s:%s:%s",
		constants.CacheKeyAccessPrefix, req.SubjectType, req.SubjectID, req.Resource, req.Action,
	)

	cached, err := s.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		return &dto.AccessCheckResponse{Allowed: cached == "1"}, nil
	}

	allowed, err := s.repo.CheckAccess(ctx, req.SubjectType, req.SubjectID, req.Resource, req.Action)
	if err != nil {
		return nil, fmt.Errorf("check access: %w", err)
	}

	val := "0"
	if allowed {
		val = "1"
	}
	s.cache.Set(ctx, cacheKey, val, time.Duration(constants.CacheTTLSeconds)*time.Second)

	return &dto.AccessCheckResponse{Allowed: allowed}, nil
}

func (s *policyService) invalidateAccessCache(ctx context.Context, subjectType, subjectID string) {
	pattern := fmt.Sprintf("%s%s:%s:*", constants.CacheKeyAccessPrefix, subjectType, subjectID)
	iter := s.cache.Scan(ctx, 0, pattern, 100).Iterator()
	for iter.Next(ctx) {
		s.cache.Del(ctx, iter.Val())
	}
}

func toRoleResponse(r *model.Role) *dto.RoleResponse {
	return &dto.RoleResponse{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}
