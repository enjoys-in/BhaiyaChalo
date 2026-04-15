package ports

import "context"

type EventPublisher interface {
	PublishUserLoggedIn(ctx context.Context, userID string, role string) error
	PublishUserLoggedOut(ctx context.Context, userID string) error
}
