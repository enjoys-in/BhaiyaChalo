package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/referral-service/internal/dto"
)

type ReferralService interface {
	GenerateCode(ctx context.Context, req *dto.GenerateCodeRequest) (*dto.ReferralCodeResponse, error)
	ApplyReferral(ctx context.Context, req *dto.ApplyReferralRequest) (*dto.ReferralResponse, error)
	CompleteReferral(ctx context.Context, referralID string) (*dto.ReferralResponse, error)
	ClaimReward(ctx context.Context, req *dto.ClaimRewardRequest) (*dto.RewardResponse, error)
	GetStats(ctx context.Context, userID string) (*dto.ReferralStatsResponse, error)
}
