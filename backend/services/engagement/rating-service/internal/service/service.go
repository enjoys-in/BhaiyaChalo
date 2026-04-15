package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/rating-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/engagement/rating-service/internal/ports"
)

type ratingService struct {
	repo      ports.RatingRepository
	publisher ports.EventPublisher
}

func NewRatingService(repo ports.RatingRepository, publisher ports.EventPublisher) ports.RatingService {
	return &ratingService{repo: repo, publisher: publisher}
}

func (s *ratingService) Create(ctx context.Context, req dto.CreateRatingRequest) (*dto.RatingResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *ratingService) GetByID(ctx context.Context, id string) (*dto.RatingResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *ratingService) Update(ctx context.Context, req dto.UpdateRatingRequest) (*dto.RatingResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *ratingService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
