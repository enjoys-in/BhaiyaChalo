package constants

const (
	ServiceName = "iam-policy-service"

	// Event types
	EventRoleCreated   = "role.created"
	EventRoleUpdated   = "role.updated"
	EventRoleDeleted   = "role.deleted"
	EventPolicyChanged = "policy.changed"
	EventPolicyCreated = "policy.created"
	EventPolicyDeleted = "policy.deleted"

	// Kafka topics
	TopicIAMEvents = "iam.events"

	// Cache keys
	CacheKeyAccessPrefix = "iam:access:"
	CacheTTLSeconds      = 300
)
