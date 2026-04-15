package dto

import "time"

type SubscriptionResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Channel   string    `json:"channel"`
	Topic     string    `json:"topic"`
	CreatedAt time.Time `json:"created_at"`
}

type SubscriptionListResponse struct {
	Subscriptions []SubscriptionResponse `json:"subscriptions"`
	Total         int64                  `json:"total"`
}
