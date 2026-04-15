package service

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/matching-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/matching-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/matching-service/internal/ports"
)

type matchingService struct {
	repo      ports.MatchRepository
	geo       ports.GeoIndex
	publisher ports.EventPublisher
}

func NewMatchingService(
	repo ports.MatchRepository,
	geo ports.GeoIndex,
	publisher ports.EventPublisher,
) ports.MatchingService {
	return &matchingService{
		repo:      repo,
		geo:       geo,
		publisher: publisher,
	}
}

func (s *matchingService) FindNearestDrivers(ctx context.Context, req *dto.FindDriversRequest) (*dto.CandidatesResponse, error) {
	radiusKM := req.RadiusKM
	if radiusKM <= 0 {
		radiusKM = 5.0
	}

	matchReq := &model.MatchRequest{
		ID:          uuid.New().String(),
		BookingID:   req.BookingID,
		CityID:      req.CityID,
		PickupLat:   req.PickupLat,
		PickupLng:   req.PickupLng,
		VehicleType: req.VehicleType,
		RadiusKM:    radiusKM,
		Status:      model.MatchSearching,
		CreatedAt:   time.Now().UTC(),
	}

	if err := s.repo.SaveRequest(ctx, matchReq); err != nil {
		return nil, fmt.Errorf("save match request: %w", err)
	}

	drivers, err := s.geo.FindNearby(ctx, req.PickupLat, req.PickupLng, radiusKM)
	if err != nil {
		return nil, fmt.Errorf("find nearby drivers: %w", err)
	}

	candidates := make([]dto.MatchResponse, 0, len(drivers))
	for _, d := range drivers {
		etaSeconds := int(d.DistanceKM * 120) // rough estimate: 2 min per km
		candidates = append(candidates, dto.MatchResponse{
			DriverID: d.DriverID,
			Distance: d.DistanceKM,
			ETA:      etaSeconds,
		})
	}

	return &dto.CandidatesResponse{
		BookingID:  req.BookingID,
		Candidates: candidates,
	}, nil
}

func (s *matchingService) AssignBestDriver(ctx context.Context, req *dto.FindDriversRequest) (*model.MatchResult, error) {
	candidatesResp, err := s.FindNearestDrivers(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(candidatesResp.Candidates) == 0 {
		_ = s.publisher.PublishMatchFailed(ctx, req.BookingID)
		return nil, fmt.Errorf("no drivers available for booking %s", req.BookingID)
	}

	scored := scoreAndRank(candidatesResp.Candidates)

	best := scored[0]
	result := &model.MatchResult{
		RequestID:  req.BookingID,
		DriverID:   best.DriverID,
		DistanceKM: best.Distance,
		ETASeconds: best.ETA,
		Score:      best.score,
	}

	mr, _ := s.repo.FindByBookingID(ctx, req.BookingID)
	if mr != nil {
		_ = s.repo.UpdateStatus(ctx, mr.ID, model.MatchMatched)
	}

	_ = s.publisher.PublishDriverMatched(ctx, result)

	return result, nil
}

type scoredCandidate struct {
	dto.MatchResponse
	score float64
}

func scoreAndRank(candidates []dto.MatchResponse) []scoredCandidate {
	scored := make([]scoredCandidate, 0, len(candidates))
	for _, c := range candidates {
		// Lower distance → higher score
		score := 1.0 / (1.0 + c.Distance)
		scored = append(scored, scoredCandidate{
			MatchResponse: c,
			score:         score,
		})
	}

	sort.Slice(scored, func(i, j int) bool {
		return scored[i].score > scored[j].score
	})

	return scored
}
