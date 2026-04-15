package dto

type SendMessageRequest struct {
	Channel       string   `json:"channel" validate:"required"`
	Payload       string   `json:"payload" validate:"required"`
	TargetUserIDs []string `json:"target_user_ids" validate:"required,min=1"`
	Priority      int      `json:"priority" validate:"gte=0,lte=2"`
}

type BroadcastRequest struct {
	Channel  string `json:"channel" validate:"required"`
	Payload  string `json:"payload" validate:"required"`
	Priority int    `json:"priority" validate:"gte=0,lte=2"`
}
