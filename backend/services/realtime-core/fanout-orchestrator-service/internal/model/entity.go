package model

import "time"

type Priority int

const (
	PriorityLow    Priority = 0
	PriorityNormal Priority = 1
	PriorityHigh   Priority = 2
)

type FanoutMessage struct {
	ID            string    `json:"id"`
	Channel       string    `json:"channel"`
	Payload       string    `json:"payload"`
	TargetUserIDs []string  `json:"target_user_ids"`
	Priority      Priority  `json:"priority"`
	CreatedAt     time.Time `json:"created_at"`
}

type DeliveryStatus struct {
	MessageID   string     `json:"message_id"`
	UserID      string     `json:"user_id"`
	Delivered   bool       `json:"delivered"`
	DeliveredAt *time.Time `json:"delivered_at,omitempty"`
}
