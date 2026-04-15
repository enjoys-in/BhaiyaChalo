package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/payment-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/payment-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/payment-service/internal/ports"
)

type KafkaWriter interface {
	WriteMessage(ctx context.Context, topic string, key string, value []byte) error
}

type Producer struct {
	writer KafkaWriter
}

func NewProducer(writer KafkaWriter) ports.EventPublisher {
	return &Producer{writer: writer}
}

type paymentEvent struct {
	PaymentID string  `json:"payment_id"`
	BookingID string  `json:"booking_id"`
	UserID    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	Method    string  `json:"method"`
	Status    string  `json:"status"`
	GatewayID string  `json:"gateway_id"`
	Timestamp string  `json:"timestamp"`
}

type refundEvent struct {
	RefundID  string  `json:"refund_id"`
	PaymentID string  `json:"payment_id"`
	Amount    float64 `json:"amount"`
	Reason    string  `json:"reason"`
	Status    string  `json:"status"`
	Timestamp string  `json:"timestamp"`
}

func (p *Producer) PublishPaymentInitiated(ctx context.Context, payment *model.Payment) error {
	return p.publishPayment(ctx, constants.TopicPaymentInitiated, payment)
}

func (p *Producer) PublishPaymentCaptured(ctx context.Context, payment *model.Payment) error {
	return p.publishPayment(ctx, constants.TopicPaymentCaptured, payment)
}

func (p *Producer) PublishPaymentFailed(ctx context.Context, payment *model.Payment) error {
	return p.publishPayment(ctx, constants.TopicPaymentFailed, payment)
}

func (p *Producer) PublishRefundIssued(ctx context.Context, refund *model.Refund) error {
	evt := refundEvent{
		RefundID:  refund.ID,
		PaymentID: refund.PaymentID,
		Amount:    refund.Amount,
		Reason:    refund.Reason,
		Status:    string(refund.Status),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	return p.writer.WriteMessage(ctx, constants.TopicRefundIssued, refund.PaymentID, data)
}

func (p *Producer) publishPayment(ctx context.Context, topic string, payment *model.Payment) error {
	evt := paymentEvent{
		PaymentID: payment.ID,
		BookingID: payment.BookingID,
		UserID:    payment.UserID,
		Amount:    payment.Amount,
		Currency:  payment.Currency,
		Method:    string(payment.Method),
		Status:    string(payment.Status),
		GatewayID: payment.GatewayID,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	return p.writer.WriteMessage(ctx, topic, payment.ID, data)
}
