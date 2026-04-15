package service

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/eta-service/internal/config"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/eta-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/eta-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/eta-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/eta-service/internal/ports"
)

type etaService struct {
	repo   ports.ETARepository
	router ports.RoutingEngine
	cfg    *config.Config
}

func NewETAService(
	repo ports.ETARepository,
	router ports.RoutingEngine,
	cfg *config.Config,
) ports.ETAService {
	return &etaService{
		repo:   repo,
		router: router,
		cfg:    cfg,
	}
}

func (s *etaService) Calculate(ctx context.Context, req *dto.CalculateETARequest) (*model.ETAResult, error) {
	etaReq := &model.ETARequest{
		FromLat:     req.FromLat,
		FromLng:     req.FromLng,
		ToLat:       req.ToLat,
		ToLng:       req.ToLng,
		VehicleType: req.VehicleType,
		CityID:      req.CityID,
	}

	cached, err := s.repo.FindRecent(ctx, etaReq)
	if err == nil && cached != nil {
		return cached, nil
	}

	route, err := s.router.GetRoute(ctx, req.FromLat, req.FromLng, req.ToLat, req.ToLng)
	if err != nil {
		return s.fallbackCalculation(ctx, etaReq)
	}

	multiplier := s.getTrafficMultiplier(req.CityID)

	result := &model.ETAResult{
		DistanceKM:        math.Round(route.DistanceKM*100) / 100,
		DurationMinutes:   math.Round(route.DurationMinutes*multiplier*100) / 100,
		TrafficMultiplier: multiplier,
		CalculatedAt:      time.Now().UTC(),
	}

	if err := s.repo.SaveCalculation(ctx, etaReq, result); err != nil {
		return result, nil
	}

	return result, nil
}

func (s *etaService) fallbackCalculation(_ context.Context, req *model.ETARequest) (*model.ETAResult, error) {
	distKM := haversineDistance(req.FromLat, req.FromLng, req.ToLat, req.ToLng)
	if distKM <= 0 {
		return nil, fmt.Errorf("invalid coordinates: distance is zero")
	}

	durationMin := (distKM / s.cfg.DefaultSpeedKMH) * 60

	result := &model.ETAResult{
		DistanceKM:        math.Round(distKM*100) / 100,
		DurationMinutes:   math.Round(durationMin*100) / 100,
		TrafficMultiplier: constants.DefaultTrafficMultiplier,
		CalculatedAt:      time.Now().UTC(),
	}
	return result, nil
}

func (s *etaService) getTrafficMultiplier(_ string) float64 {
	return constants.DefaultTrafficMultiplier
}

func haversineDistance(lat1, lng1, lat2, lng2 float64) float64 {
	const earthRadiusKM = 6371.0

	dLat := degreesToRadians(lat2 - lat1)
	dLng := degreesToRadians(lng2 - lng1)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(degreesToRadians(lat1))*math.Cos(degreesToRadians(lat2))*
			math.Sin(dLng/2)*math.Sin(dLng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusKM * c
}

func degreesToRadians(deg float64) float64 {
	return deg * math.Pi / 180
}
