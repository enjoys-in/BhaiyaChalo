package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/review-moderation-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/engagement/review-moderation-service/internal/ports"
)

type reviewModerationService struct {
	repo      ports.ReviewModerationRepository
	publisher ports.EventPublisher
}

func NewReviewModerationService(repo ports.ReviewModerationRepository, publisher ports.EventPublisher) ports.ReviewModerationService {
	return &reviewModerationService{repo: repo, publisher: publisher}
}

func (s *reviewModerationService) Create(ctx context.Context, req dto.CreateReviewRequest) (*dto.ReviewResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *reviewModerationService) GetByID(ctx context.Context, id string) (*dto.ReviewResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *reviewModerationService) Update(ctx context.Context, req dto.UpdateReviewRequest) (*dto.ReviewResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *reviewModerationService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
