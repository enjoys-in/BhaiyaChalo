package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/earnings-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/earnings-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/earnings-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/earnings-service/internal/ports"
)

type earningsService struct {
	repo      ports.EarningsRepository
	publisher ports.EventPublisher
}

func NewEarningsService(repo ports.EarningsRepository, publisher ports.EventPublisher) ports.EarningsService {
	return &earningsService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *earningsService) RecordEarning(ctx context.Context, req dto.RecordEarningRequest) (*dto.EarningResponse, error) {
	now := time.Now().UTC()

	commission := req.FareAmount * constants.CommissionRate
	netEarning := req.FareAmount - commission + req.IncentiveBonus + req.TipAmount

	earning := &model.Earning{
		ID:             uuid.NewString(),
		DriverID:       req.DriverID,
		TripID:         req.TripID,
		FareAmount:     req.FareAmount,
		Commission:     commission,
		IncentiveBonus: req.IncentiveBonus,
		TipAmount:      req.TipAmount,
		NetEarning:     netEarning,
		Currency:       req.Currency,
		EarnedAt:       now,
		CreatedAt:      now,
	}

	if err := s.repo.Record(ctx, earning); err != nil {
		return nil, fmt.Errorf("record earning: %w", err)
	}

	_ = s.publisher.PublishEarningRecorded(ctx, earning.DriverID, earning.TripID, earning.NetEarning)

	return toEarningResponse(earning), nil
}

func (s *earningsService) GetEarnings(ctx context.Context, driverID string, from, to time.Time) ([]dto.EarningResponse, error) {
	earnings, err := s.repo.FindByDriverID(ctx, driverID, from, to)
	if err != nil {
		return nil, fmt.Errorf("find earnings: %w", err)
	}

	resp := make([]dto.EarningResponse, len(earnings))
	for i, e := range earnings {
		resp[i] = *toEarningResponse(e)
	}
	return resp, nil
}

func (s *earningsService) GetDailySummary(ctx context.Context, driverID string, date time.Time) (*dto.DailySummaryResponse, error) {
	summary, err := s.repo.GetDailySummary(ctx, driverID, date)
	if err != nil {
		return nil, fmt.Errorf("get daily summary: %w", err)
	}

	return &dto.DailySummaryResponse{
		DriverID:        summary.DriverID,
		Date:            summary.Date,
		TotalTrips:      summary.TotalTrips,
		TotalFare:       summary.TotalFare,
		TotalCommission: summary.TotalCommission,
		TotalIncentive:  summary.TotalIncentive,
		TotalTips:       summary.TotalTips,
		NetEarning:      summary.NetEarning,
	}, nil
}

func (s *earningsService) GetWeeklySummary(ctx context.Context, driverID string, weekStart time.Time) (*dto.WeeklySummaryResponse, error) {
	summary, err := s.repo.GetWeeklySummary(ctx, driverID, weekStart)
	if err != nil {
		return nil, fmt.Errorf("get weekly summary: %w", err)
	}

	return &dto.WeeklySummaryResponse{
		DriverID:        summary.DriverID,
		WeekStart:       summary.WeekStart,
		WeekEnd:         summary.WeekEnd,
		TotalTrips:      summary.TotalTrips,
		TotalFare:       summary.TotalFare,
		TotalCommission: summary.TotalCommission,
		TotalIncentive:  summary.TotalIncentive,
		TotalTips:       summary.TotalTips,
		NetEarning:      summary.NetEarning,
	}, nil
}

func toEarningResponse(e *model.Earning) *dto.EarningResponse {
	return &dto.EarningResponse{
		ID:             e.ID,
		DriverID:       e.DriverID,
		TripID:         e.TripID,
		FareAmount:     e.FareAmount,
		Commission:     e.Commission,
		IncentiveBonus: e.IncentiveBonus,
		TipAmount:      e.TipAmount,
		NetEarning:     e.NetEarning,
		Currency:       e.Currency,
		EarnedAt:       e.EarnedAt,
	}
}
