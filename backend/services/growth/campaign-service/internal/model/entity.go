package model

import "time"

type CampaignType string

const (
	CampaignTypePush  CampaignType = "push"
	CampaignTypeSMS   CampaignType = "sms"
	CampaignTypeEmail CampaignType = "email"
	CampaignTypeInApp CampaignType = "in_app"
)

type CampaignStatus string

const (
	CampaignStatusDraft     CampaignStatus = "draft"
	CampaignStatusScheduled CampaignStatus = "scheduled"
	CampaignStatusActive    CampaignStatus = "active"
	CampaignStatusCompleted CampaignStatus = "completed"
	CampaignStatusPaused    CampaignStatus = "paused"
)

type Campaign struct {
	ID             string         `json:"id" db:"id"`
	Name           string         `json:"name" db:"name"`
	Description    string         `json:"description" db:"description"`
	Type           CampaignType   `json:"type" db:"type"`
	TargetAudience string         `json:"target_audience" db:"target_audience"`
	CityID         string         `json:"city_id" db:"city_id"`
	PromoCodeID    string         `json:"promo_code_id" db:"promo_code_id"`
	Status         CampaignStatus `json:"status" db:"status"`
	ScheduledAt    *time.Time     `json:"scheduled_at,omitempty" db:"scheduled_at"`
	StartedAt      *time.Time     `json:"started_at,omitempty" db:"started_at"`
	CompletedAt    *time.Time     `json:"completed_at,omitempty" db:"completed_at"`
	CreatedAt      time.Time      `json:"created_at" db:"created_at"`
}

type CampaignStats struct {
	CampaignID string `json:"campaign_id" db:"campaign_id"`
	Sent       int    `json:"sent" db:"sent"`
	Delivered  int    `json:"delivered" db:"delivered"`
	Opened     int    `json:"opened" db:"opened"`
	Clicked    int    `json:"clicked" db:"clicked"`
	Converted  int    `json:"converted" db:"converted"`
}
