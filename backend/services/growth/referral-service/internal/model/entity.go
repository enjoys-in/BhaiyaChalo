package model

import "time"

type ReferralStatus string

const (
	ReferralStatusPending   ReferralStatus = "pending"
	ReferralStatusCompleted ReferralStatus = "completed"
	ReferralStatusRewarded  ReferralStatus = "rewarded"
)

type RewardType string

const (
	RewardTypeReferrer RewardType = "referrer"
	RewardTypeReferee  RewardType = "referee"
)

type ReferralCode struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Code      string    `json:"code" db:"code"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Referral struct {
	ID             string         `json:"id" db:"id"`
	ReferrerID     string         `json:"referrer_id" db:"referrer_id"`
	RefereeID      string         `json:"referee_id" db:"referee_id"`
	ReferralCodeID string         `json:"referral_code_id" db:"referral_code_id"`
	Status         ReferralStatus `json:"status" db:"status"`
	CompletedAt    *time.Time     `json:"completed_at,omitempty" db:"completed_at"`
	CreatedAt      time.Time      `json:"created_at" db:"created_at"`
}

type ReferralReward struct {
	ID         string     `json:"id" db:"id"`
	ReferralID string     `json:"referral_id" db:"referral_id"`
	UserID     string     `json:"user_id" db:"user_id"`
	Type       RewardType `json:"type" db:"type"`
	Amount     float64    `json:"amount" db:"amount"`
	Currency   string     `json:"currency" db:"currency"`
	CreditedAt *time.Time `json:"credited_at,omitempty" db:"credited_at"`
}
