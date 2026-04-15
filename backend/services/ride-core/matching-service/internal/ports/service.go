package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/matching-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/matching-service/internal/model"
)

type MatchingService interface {
	FindNearestDrivers(ctx context.Context, req *dto.FindDriversRequest) (*dto.CandidatesResponse, error)
	AssignBestDriver(ctx context.Context, req *dto.FindDriversRequest) (*model.MatchResult, error)
}
