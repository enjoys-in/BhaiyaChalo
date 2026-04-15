package ports

import "context"

type EventPublisher interface {
	PublishMessageDelivered(ctx context.Context, messageID string, userID string) error
	PublishMessageFailed(ctx context.Context, messageID string, userID string, reason string) error
}
