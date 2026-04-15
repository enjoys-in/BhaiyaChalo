package dto

import "time"

type CampaignResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Type           string `json:"type"`
	TargetAudience string `json:"target_audience"`
	CityID         string `json:"city_id"`
	PromoCodeID    string `json:"promo_code_id,omitempty"`
	Status         string `json:"status"`
	ScheduledAt    string `json:"scheduled_at,omitempty"`
	StartedAt      string `json:"started_at,omitempty"`
	CompletedAt    string `json:"completed_at,omitempty"`
	CreatedAt      string `json:"created_at"`
}

type CampaignStatsResponse struct {
	CampaignID string `json:"campaign_id"`
	Sent       int    `json:"sent"`
	Delivered  int    `json:"delivered"`
	Opened     int    `json:"opened"`
	Clicked    int    `json:"clicked"`
	Converted  int    `json:"converted"`
}

type CampaignListResponse struct {
	Campaigns []CampaignResponse `json:"campaigns"`
	Total     int                `json:"total"`
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
