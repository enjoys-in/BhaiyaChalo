package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/referral-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/referral-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/referral-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/referral-service/internal/ports"
	"github.com/google/uuid"
)

type referralService struct {
	repo      ports.ReferralRepository
	publisher ports.ReferralEventPublisher
}

func NewReferralService(repo ports.ReferralRepository, publisher ports.ReferralEventPublisher) ports.ReferralService {
	return &referralService{repo: repo, publisher: publisher}
}

func (s *referralService) GenerateCode(ctx context.Context, req *dto.GenerateCodeRequest) (*dto.ReferralCodeResponse, error) {
	existing, err := s.repo.FindByUserID(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("checking existing code: %w", err)
	}
	if existing != nil {
		return toReferralCodeResponse(existing), nil
	}

	code, err := generateUniqueCode()
	if err != nil {
		return nil, fmt.Errorf("generating code: %w", err)
	}

	rc := &model.ReferralCode{
		ID:        uuid.New().String(),
		UserID:    req.UserID,
		Code:      code,
		Active:    true,
		CreatedAt: time.Now(),
	}

	if err := s.repo.CreateCode(ctx, rc); err != nil {
		return nil, fmt.Errorf("creating referral code: %w", err)
	}

	return toReferralCodeResponse(rc), nil
}

func (s *referralService) ApplyReferral(ctx context.Context, req *dto.ApplyReferralRequest) (*dto.ReferralResponse, error) {
	rc, err := s.repo.FindByCode(ctx, strings.ToUpper(req.Code))
	if err != nil {
		return nil, fmt.Errorf("finding referral code: %w", err)
	}
	if rc == nil {
		return nil, fmt.Errorf("referral code not found")
	}
	if !rc.Active {
		return nil, fmt.Errorf("referral code is inactive")
	}
	if rc.UserID == req.RefereeID {
		return nil, fmt.Errorf("cannot use own referral code")
	}

	referral := &model.Referral{
		ID:             uuid.New().String(),
		ReferrerID:     rc.UserID,
		RefereeID:      req.RefereeID,
		ReferralCodeID: rc.ID,
		Status:         model.ReferralStatusPending,
		CreatedAt:      time.Now(),
	}

	if err := s.repo.CreateReferral(ctx, referral); err != nil {
		return nil, fmt.Errorf("creating referral: %w", err)
	}

	_ = s.publisher.PublishReferralApplied(ctx, referral)

	return toReferralResponse(referral), nil
}

func (s *referralService) CompleteReferral(ctx context.Context, referralID string) (*dto.ReferralResponse, error) {
	referral, err := s.repo.FindReferralByID(ctx, referralID)
	if err != nil {
		return nil, fmt.Errorf("finding referral: %w", err)
	}
	if referral == nil {
		return nil, fmt.Errorf("referral not found")
	}
	if referral.Status != model.ReferralStatusPending {
		return nil, fmt.Errorf("referral is not in pending status")
	}

	now := time.Now()
	referral.Status = model.ReferralStatusCompleted
	referral.CompletedAt = &now

	if err := s.repo.UpdateReferralStatus(ctx, referralID, model.ReferralStatusCompleted); err != nil {
		return nil, fmt.Errorf("updating referral: %w", err)
	}

	_ = s.publisher.PublishReferralCompleted(ctx, referral)

	return toReferralResponse(referral), nil
}

func (s *referralService) ClaimReward(ctx context.Context, req *dto.ClaimRewardRequest) (*dto.RewardResponse, error) {
	referral, err := s.repo.FindReferralByID(ctx, req.ReferralID)
	if err != nil {
		return nil, fmt.Errorf("finding referral: %w", err)
	}
	if referral == nil {
		return nil, fmt.Errorf("referral not found")
	}
	if referral.Status != model.ReferralStatusCompleted {
		return nil, fmt.Errorf("referral must be completed before claiming reward")
	}

	rewardType := model.RewardTypeReferrer
	if req.UserID == referral.RefereeID {
		rewardType = model.RewardTypeReferee
	}

	now := time.Now()
	reward := &model.ReferralReward{
		ID:         uuid.New().String(),
		ReferralID: req.ReferralID,
		UserID:     req.UserID,
		Type:       rewardType,
		Amount:     constants.DefaultRewardAmount,
		Currency:   constants.DefaultCurrency,
		CreditedAt: &now,
	}

	if err := s.repo.CreateReward(ctx, reward); err != nil {
		return nil, fmt.Errorf("creating reward: %w", err)
	}

	_ = s.repo.UpdateReferralStatus(ctx, req.ReferralID, model.ReferralStatusRewarded)
	_ = s.publisher.PublishRewardCredited(ctx, reward)

	return toRewardResponse(reward), nil
}

func (s *referralService) GetStats(ctx context.Context, userID string) (*dto.ReferralStatsResponse, error) {
	total, completed, pending, earnings, err := s.repo.GetStats(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("getting stats: %w", err)
	}

	return &dto.ReferralStatsResponse{
		UserID:         userID,
		TotalReferrals: total,
		Completed:      completed,
		Pending:        pending,
		TotalEarnings:  earnings,
	}, nil
}

func generateUniqueCode() (string, error) {
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return strings.ToUpper("REF" + hex.EncodeToString(b)), nil
}

func toReferralCodeResponse(rc *model.ReferralCode) *dto.ReferralCodeResponse {
	return &dto.ReferralCodeResponse{
		ID:        rc.ID,
		UserID:    rc.UserID,
		Code:      rc.Code,
		Active:    rc.Active,
		CreatedAt: dto.FormatTime(rc.CreatedAt),
	}
}

func toReferralResponse(r *model.Referral) *dto.ReferralResponse {
	return &dto.ReferralResponse{
		ID:             r.ID,
		ReferrerID:     r.ReferrerID,
		RefereeID:      r.RefereeID,
		ReferralCodeID: r.ReferralCodeID,
		Status:         string(r.Status),
		CompletedAt:    dto.FormatTimePtr(r.CompletedAt),
		CreatedAt:      dto.FormatTime(r.CreatedAt),
	}
}

func toRewardResponse(rw *model.ReferralReward) *dto.RewardResponse {
	return &dto.RewardResponse{
		ID:         rw.ID,
		ReferralID: rw.ReferralID,
		UserID:     rw.UserID,
		Type:       string(rw.Type),
		Amount:     rw.Amount,
		Currency:   rw.Currency,
		CreditedAt: dto.FormatTimePtr(rw.CreditedAt),
	}
}
