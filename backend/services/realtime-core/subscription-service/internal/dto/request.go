package dto

type SubscribeRequest struct {
	UserID  string `json:"user_id" validate:"required"`
	Channel string `json:"channel" validate:"required"`
	Topic   string `json:"topic" validate:"required"`
}

type UnsubscribeRequest struct {
	UserID         string `json:"user_id" validate:"required"`
	SubscriptionID string `json:"subscription_id" validate:"required"`
}
