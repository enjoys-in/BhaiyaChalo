package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/review-moderation-service/internal/model"
)

type EventPublisher interface {
	PublishReviewSubmitted(ctx context.Context, entity *model.Review) error
	PublishReviewModerated(ctx context.Context, entity *model.Review) error
}
