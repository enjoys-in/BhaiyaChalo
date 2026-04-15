package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/promo-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/promo-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/promo-service/internal/ports"
	"github.com/google/uuid"
)

type promoService struct {
	repo      ports.PromoRepository
	publisher ports.PromoEventPublisher
}

func NewPromoService(repo ports.PromoRepository, publisher ports.PromoEventPublisher) ports.PromoService {
	return &promoService{repo: repo, publisher: publisher}
}

func (s *promoService) Create(ctx context.Context, req *dto.CreatePromoRequest) (*dto.PromoResponse, error) {
	validFrom, validUntil, err := req.ParseTimes()
	if err != nil {
		return nil, fmt.Errorf("invalid time format: %w", err)
	}

	existing, err := s.repo.FindByCode(ctx, strings.ToUpper(req.Code))
	if err != nil {
		return nil, fmt.Errorf("checking existing code: %w", err)
	}
	if existing != nil {
		return nil, fmt.Errorf("promo code already exists")
	}

	promo := &model.PromoCode{
		ID:            uuid.New().String(),
		Code:          strings.ToUpper(req.Code),
		CityID:        req.CityID,
		Type:          model.PromoType(req.Type),
		DiscountValue: req.DiscountValue,
		MaxDiscount:   req.MaxDiscount,
		MinOrderValue: req.MinOrderValue,
		UsageLimit:    req.UsageLimit,
		UsedCount:     0,
		ValidFrom:     validFrom,
		ValidUntil:    validUntil,
		Active:        true,
		CreatedAt:     time.Now(),
	}

	if err := s.repo.Create(ctx, promo); err != nil {
		return nil, fmt.Errorf("creating promo: %w", err)
	}

	_ = s.publisher.PublishPromoCreated(ctx, promo)

	return toPromoResponse(promo), nil
}

func (s *promoService) Apply(ctx context.Context, req *dto.ApplyPromoRequest) (*dto.ApplyPromoResponse, error) {
	promo, err := s.repo.FindByCode(ctx, strings.ToUpper(req.Code))
	if err != nil {
		return nil, fmt.Errorf("finding promo: %w", err)
	}

	if resp := s.checkEligibility(ctx, promo, req.UserID, req.CityID, req.BookingAmount); resp != nil {
		return resp, nil
	}

	discount := s.calculateDiscount(promo, req.BookingAmount)

	usage := &model.PromoUsage{
		ID:              uuid.New().String(),
		PromoID:         promo.ID,
		UserID:          req.UserID,
		BookingID:       "",
		DiscountApplied: discount,
		UsedAt:          time.Now(),
	}

	if err := s.repo.IncrementUsage(ctx, promo.ID); err != nil {
		return nil, fmt.Errorf("incrementing usage: %w", err)
	}
	if err := s.repo.RecordUsage(ctx, usage); err != nil {
		return nil, fmt.Errorf("recording usage: %w", err)
	}

	_ = s.publisher.PublishPromoApplied(ctx, usage)

	return &dto.ApplyPromoResponse{
		Valid:          true,
		DiscountAmount: discount,
		FinalAmount:    req.BookingAmount - discount,
	}, nil
}

func (s *promoService) Validate(ctx context.Context, req *dto.ValidatePromoRequest) (*dto.ApplyPromoResponse, error) {
	promo, err := s.repo.FindByCode(ctx, strings.ToUpper(req.Code))
	if err != nil {
		return nil, fmt.Errorf("finding promo: %w", err)
	}

	if resp := s.checkEligibility(ctx, promo, req.UserID, req.CityID, 0); resp != nil {
		return resp, nil
	}

	return &dto.ApplyPromoResponse{Valid: true, Message: "promo code is valid"}, nil
}

func (s *promoService) ListActive(ctx context.Context, cityID string) (*dto.PromoListResponse, error) {
	var promos []model.PromoCode
	var err error

	if cityID != "" {
		promos, err = s.repo.ListByCityID(ctx, cityID)
	} else {
		promos, err = s.repo.ListActive(ctx)
	}
	if err != nil {
		return nil, fmt.Errorf("listing promos: %w", err)
	}

	responses := make([]dto.PromoResponse, 0, len(promos))
	for i := range promos {
		responses = append(responses, *toPromoResponse(&promos[i]))
	}

	return &dto.PromoListResponse{Promos: responses, Total: len(responses)}, nil
}

func (s *promoService) checkEligibility(ctx context.Context, promo *model.PromoCode, userID, cityID string, amount float64) *dto.ApplyPromoResponse {
	if promo == nil {
		return &dto.ApplyPromoResponse{Valid: false, Message: "promo code not found"}
	}
	if !promo.Active {
		return &dto.ApplyPromoResponse{Valid: false, Message: "promo code is inactive"}
	}
	now := time.Now()
	if now.Before(promo.ValidFrom) || now.After(promo.ValidUntil) {
		return &dto.ApplyPromoResponse{Valid: false, Message: "promo code has expired"}
	}
	if promo.UsedCount >= promo.UsageLimit {
		return &dto.ApplyPromoResponse{Valid: false, Message: "promo code usage limit reached"}
	}
	if promo.CityID != cityID {
		return &dto.ApplyPromoResponse{Valid: false, Message: "promo code not valid for this city"}
	}
	if amount > 0 && amount < promo.MinOrderValue {
		return &dto.ApplyPromoResponse{Valid: false, Message: "minimum order value not met"}
	}

	used, err := s.repo.HasUserUsed(ctx, promo.ID, userID)
	if err == nil && used {
		return &dto.ApplyPromoResponse{Valid: false, Message: "user has already used this promo code"}
	}

	return nil
}

func (s *promoService) calculateDiscount(promo *model.PromoCode, amount float64) float64 {
	var discount float64
	switch promo.Type {
	case model.PromoTypeFlat:
		discount = promo.DiscountValue
	case model.PromoTypePercentage:
		discount = amount * promo.DiscountValue / 100
	}
	if discount > promo.MaxDiscount {
		discount = promo.MaxDiscount
	}
	if discount > amount {
		discount = amount
	}
	return discount
}

func toPromoResponse(p *model.PromoCode) *dto.PromoResponse {
	return &dto.PromoResponse{
		ID:            p.ID,
		Code:          p.Code,
		CityID:        p.CityID,
		Type:          string(p.Type),
		DiscountValue: p.DiscountValue,
		MaxDiscount:   p.MaxDiscount,
		MinOrderValue: p.MinOrderValue,
		UsageLimit:    p.UsageLimit,
		UsedCount:     p.UsedCount,
		ValidFrom:     dto.FormatTime(p.ValidFrom),
		ValidUntil:    dto.FormatTime(p.ValidUntil),
		Active:        p.Active,
		CreatedAt:     dto.FormatTime(p.CreatedAt),
	}
}
