package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/referral-service/internal/model"
)

type ReferralEventPublisher interface {
	PublishReferralApplied(ctx context.Context, referral *model.Referral) error
	PublishReferralCompleted(ctx context.Context, referral *model.Referral) error
	PublishRewardCredited(ctx context.Context, reward *model.ReferralReward) error
}
