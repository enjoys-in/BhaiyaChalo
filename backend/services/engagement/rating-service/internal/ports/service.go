package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/rating-service/internal/dto"
)

type RatingService interface {
	Create(ctx context.Context, req dto.CreateRatingRequest) (*dto.RatingResponse, error)
	GetByID(ctx context.Context, id string) (*dto.RatingResponse, error)
	Update(ctx context.Context, req dto.UpdateRatingRequest) (*dto.RatingResponse, error)
	Delete(ctx context.Context, id string) error
}
