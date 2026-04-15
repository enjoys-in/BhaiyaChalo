package ports

import "context"

type EventPublisher interface {
	PublishRoleCreated(ctx context.Context, roleID, roleName string) error
	PublishPolicyChanged(ctx context.Context, policyID, subjectType, subjectID string) error
}
