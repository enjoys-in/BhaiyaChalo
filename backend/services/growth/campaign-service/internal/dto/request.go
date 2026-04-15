package dto

type CreateCampaignRequest struct {
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description"`
	Type           string `json:"type" validate:"required,oneof=push sms email in_app"`
	TargetAudience string `json:"target_audience" validate:"required"`
	CityID         string `json:"city_id" validate:"required"`
	PromoCodeID    string `json:"promo_code_id,omitempty"`
	ScheduledAt    string `json:"scheduled_at,omitempty"`
}

type UpdateCampaignRequest struct {
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	TargetAudience string `json:"target_audience,omitempty"`
	ScheduledAt    string `json:"scheduled_at,omitempty"`
}

type LaunchCampaignRequest struct {
	CampaignID string `json:"campaign_id" validate:"required"`
}
