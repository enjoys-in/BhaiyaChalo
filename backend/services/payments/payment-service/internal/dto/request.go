package dto

type InitiatePaymentRequest struct {
	BookingID string  `json:"booking_id" validate:"required"`
	UserID    string  `json:"user_id" validate:"required"`
	Amount    float64 `json:"amount" validate:"required,gt=0"`
	Currency  string  `json:"currency" validate:"required"`
	Method    string  `json:"method" validate:"required,oneof=card upi wallet cash"`
}

type CapturePaymentRequest struct {
	PaymentID string `json:"payment_id" validate:"required"`
	GatewayID string `json:"gateway_id" validate:"required"`
}

type RefundRequest struct {
	PaymentID string  `json:"payment_id" validate:"required"`
	Amount    float64 `json:"amount" validate:"required,gt=0"`
	Reason    string  `json:"reason" validate:"required"`
}
