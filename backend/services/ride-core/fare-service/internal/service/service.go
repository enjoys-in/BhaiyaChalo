package service

import (
	"context"
	"math"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/fare-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/fare-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/fare-service/internal/event"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/fare-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/fare-service/internal/ports"
)

type fareService struct {
	repo     ports.FareRepository
	producer *event.Producer
}

func NewFareService(repo ports.FareRepository, producer *event.Producer) ports.FareService {
	return &fareService{
		repo:     repo,
		producer: producer,
	}
}

func (s *fareService) Calculate(ctx context.Context, req dto.CalculateFareRequest) (*model.FareCalculation, error) {
	cfg, err := s.repo.GetConfig(ctx, req.CityID, req.VehicleType)
	if err != nil {
		return nil, err
	}

	calc := computeFare(cfg, req)
	calc.ID = uuid.New().String()
	calc.BookingID = req.BookingID
	calc.CityID = req.CityID
	calc.VehicleType = req.VehicleType
	calc.Currency = constants.DefaultCurrency
	calc.CreatedAt = time.Now().UTC()

	if err := s.repo.SaveCalculation(ctx, calc); err != nil {
		return nil, err
	}

	if err := s.producer.PublishFareCalculated(ctx, calc); err != nil {
		return nil, err
	}

	return calc, nil
}

func (s *fareService) Recalculate(ctx context.Context, req dto.RecalculateFareRequest) (*model.FareCalculation, error) {
	existing, err := s.repo.FindByBookingID(ctx, req.BookingID)
	if err != nil {
		return nil, err
	}

	calcReq := dto.CalculateFareRequest{
		BookingID:       req.BookingID,
		DistanceKM:      req.DistanceKM,
		DurationMin:     req.DurationMin,
		CityID:          existing.CityID,
		VehicleType:     existing.VehicleType,
		SurgeMultiplier: req.SurgeMultiplier,
		PromoDiscount:   req.PromoDiscount,
		TollCharges:     req.TollCharges,
	}

	return s.Calculate(ctx, calcReq)
}

func (s *fareService) GetBreakdown(ctx context.Context, bookingID string) (*model.FareCalculation, error) {
	return s.repo.FindByBookingID(ctx, bookingID)
}

func computeFare(cfg *model.FareConfig, req dto.CalculateFareRequest) *model.FareCalculation {
	distanceCharge := req.DistanceKM * cfg.BasePricePerKM
	timeCharge := req.DurationMin * cfg.BasePricePerMin
	basePrice := distanceCharge + timeCharge

	surgeMultiplier := req.SurgeMultiplier
	if surgeMultiplier < 1.0 {
		surgeMultiplier = 1.0
	}

	surgeAmount := basePrice * (surgeMultiplier - 1.0)
	subtotal := basePrice + surgeAmount + req.TollCharges
	taxAmount := subtotal * constants.TaxRate
	totalFare := subtotal + taxAmount - req.PromoDiscount

	if totalFare < cfg.MinFare {
		totalFare = cfg.MinFare
	}

	totalFare = math.Round(totalFare*100) / 100

	return &model.FareCalculation{
		BasePrice:       roundTwo(basePrice),
		DistanceCharge:  roundTwo(distanceCharge),
		TimeCharge:      roundTwo(timeCharge),
		SurgeMultiplier: surgeMultiplier,
		SurgeAmount:     roundTwo(surgeAmount),
		TollCharges:     roundTwo(req.TollCharges),
		TaxAmount:       roundTwo(taxAmount),
		PromoDiscount:   roundTwo(req.PromoDiscount),
		TotalFare:       totalFare,
	}
}

func roundTwo(v float64) float64 {
	return math.Round(v*100) / 100
}
