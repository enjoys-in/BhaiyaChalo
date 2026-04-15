package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/referral-service/internal/model"
)

type ReferralRepository interface {
	CreateCode(ctx context.Context, code *model.ReferralCode) error
	FindByCode(ctx context.Context, code string) (*model.ReferralCode, error)
	FindByUserID(ctx context.Context, userID string) (*model.ReferralCode, error)
	CreateReferral(ctx context.Context, referral *model.Referral) error
	UpdateReferralStatus(ctx context.Context, id string, status model.ReferralStatus) error
	CreateReward(ctx context.Context, reward *model.ReferralReward) error
	GetStats(ctx context.Context, userID string) (total, completed, pending int, earnings float64, err error)
	FindReferralByID(ctx context.Context, id string) (*model.Referral, error)
}
