package dto

type GenerateCodeRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

type ApplyReferralRequest struct {
	Code      string `json:"code" validate:"required"`
	RefereeID string `json:"referee_id" validate:"required"`
}

type ClaimRewardRequest struct {
	ReferralID string `json:"referral_id" validate:"required"`
	UserID     string `json:"user_id" validate:"required"`
}
