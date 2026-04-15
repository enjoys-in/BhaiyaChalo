package dto

import "time"

type ReferralCodeResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Code      string `json:"code"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"created_at"`
}

type ReferralResponse struct {
	ID             string `json:"id"`
	ReferrerID     string `json:"referrer_id"`
	RefereeID      string `json:"referee_id"`
	ReferralCodeID string `json:"referral_code_id"`
	Status         string `json:"status"`
	CompletedAt    string `json:"completed_at,omitempty"`
	CreatedAt      string `json:"created_at"`
}

type RewardResponse struct {
	ID         string  `json:"id"`
	ReferralID string  `json:"referral_id"`
	UserID     string  `json:"user_id"`
	Type       string  `json:"type"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	CreditedAt string  `json:"credited_at,omitempty"`
}

type ReferralStatsResponse struct {
	UserID         string `json:"user_id"`
	TotalReferrals int    `json:"total_referrals"`
	Completed      int    `json:"completed"`
	Pending        int    `json:"pending"`
	TotalEarnings  float64 `json:"total_earnings"`
}

func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

func FormatTimePtr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(time.RFC3339)
}
