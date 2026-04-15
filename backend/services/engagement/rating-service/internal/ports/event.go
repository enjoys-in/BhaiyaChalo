package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/rating-service/internal/model"
)

type EventPublisher interface {
	PublishRatingSubmitted(ctx context.Context, entity *model.Rating) error
	PublishRatingUpdated(ctx context.Context, entity *model.Rating) error
}
