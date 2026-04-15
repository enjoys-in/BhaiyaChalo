package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/matching-service/internal/model"
)

type EventPublisher interface {
	PublishDriverMatched(ctx context.Context, result *model.MatchResult) error
	PublishMatchFailed(ctx context.Context, bookingID string) error
}
