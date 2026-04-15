package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/review-moderation-service/internal/dto"
)

type ReviewModerationService interface {
	Create(ctx context.Context, req dto.CreateReviewRequest) (*dto.ReviewResponse, error)
	GetByID(ctx context.Context, id string) (*dto.ReviewResponse, error)
	Update(ctx context.Context, req dto.UpdateReviewRequest) (*dto.ReviewResponse, error)
	Delete(ctx context.Context, id string) error
}
