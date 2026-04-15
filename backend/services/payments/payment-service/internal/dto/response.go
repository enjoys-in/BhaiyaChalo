package dto

import "time"

type PaymentResponse struct {
	ID            string    `json:"id"`
	BookingID     string    `json:"booking_id"`
	UserID        string    `json:"user_id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Method        string    `json:"method"`
	GatewayID     string    `json:"gateway_id"`
	GatewayStatus string    `json:"gateway_status"`
	Status        string    `json:"status"`
	FailureReason string    `json:"failure_reason,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type RefundResponse struct {
	ID        string    `json:"id"`
	PaymentID string    `json:"payment_id"`
	Amount    float64   `json:"amount"`
	Reason    string    `json:"reason"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type PaymentStatusResponse struct {
	PaymentID string `json:"payment_id"`
	Status    string `json:"status"`
	GatewayID string `json:"gateway_id,omitempty"`
}
