package model

import "time"

type PaymentMethod string

const (
	PaymentMethodCard   PaymentMethod = "card"
	PaymentMethodUPI    PaymentMethod = "upi"
	PaymentMethodWallet PaymentMethod = "wallet"
	PaymentMethodCash   PaymentMethod = "cash"
)

type PaymentStatus string

const (
	PaymentStatusInitiated  PaymentStatus = "initiated"
	PaymentStatusAuthorized PaymentStatus = "authorized"
	PaymentStatusCaptured   PaymentStatus = "captured"
	PaymentStatusFailed     PaymentStatus = "failed"
	PaymentStatusRefunded   PaymentStatus = "refunded"
)

type RefundStatus string

const (
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusProcessed RefundStatus = "processed"
	RefundStatusFailed    RefundStatus = "failed"
)

type Payment struct {
	ID            string        `json:"id" db:"id"`
	BookingID     string        `json:"booking_id" db:"booking_id"`
	UserID        string        `json:"user_id" db:"user_id"`
	Amount        float64       `json:"amount" db:"amount"`
	Currency      string        `json:"currency" db:"currency"`
	Method        PaymentMethod `json:"method" db:"method"`
	GatewayID     string        `json:"gateway_id" db:"gateway_id"`
	GatewayStatus string        `json:"gateway_status" db:"gateway_status"`
	Status        PaymentStatus `json:"status" db:"status"`
	FailureReason string        `json:"failure_reason,omitempty" db:"failure_reason"`
	CreatedAt     time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at" db:"updated_at"`
}

type Refund struct {
	ID        string       `json:"id" db:"id"`
	PaymentID string       `json:"payment_id" db:"payment_id"`
	Amount    float64      `json:"amount" db:"amount"`
	Reason    string       `json:"reason" db:"reason"`
	Status    RefundStatus `json:"status" db:"status"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
}
