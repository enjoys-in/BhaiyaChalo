package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/payment-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/payment-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/payment-service/internal/ports"
)

type paymentService struct {
	repo      ports.PaymentRepository
	gateway   ports.PaymentGateway
	publisher ports.EventPublisher
}

func NewPaymentService(repo ports.PaymentRepository, gateway ports.PaymentGateway, publisher ports.EventPublisher) ports.PaymentService {
	return &paymentService{
		repo:      repo,
		gateway:   gateway,
		publisher: publisher,
	}
}

func (s *paymentService) Initiate(ctx context.Context, req dto.InitiatePaymentRequest) (*model.Payment, error) {
	now := time.Now().UTC()
	payment := &model.Payment{
		ID:        uuid.New().String(),
		BookingID: req.BookingID,
		UserID:    req.UserID,
		Amount:    req.Amount,
		Currency:  req.Currency,
		Method:    model.PaymentMethod(req.Method),
		Status:    model.PaymentStatusInitiated,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := s.gateway.Authorize(ctx, req.Amount, req.Currency, req.Method)
	if err != nil {
		payment.Status = model.PaymentStatusFailed
		payment.FailureReason = err.Error()
		_ = s.repo.Create(ctx, payment)
		_ = s.publisher.PublishPaymentFailed(ctx, payment)
		return payment, fmt.Errorf("gateway authorization failed: %w", err)
	}

	payment.GatewayID = result.GatewayID
	payment.GatewayStatus = result.GatewayStatus
	payment.Status = model.PaymentStatusAuthorized

	if err := s.repo.Create(ctx, payment); err != nil {
		return nil, err
	}

	_ = s.publisher.PublishPaymentInitiated(ctx, payment)
	return payment, nil
}

func (s *paymentService) Capture(ctx context.Context, req dto.CapturePaymentRequest) (*model.Payment, error) {
	payment, err := s.repo.FindByID(ctx, req.PaymentID)
	if err != nil {
		return nil, err
	}

	if payment.Status != model.PaymentStatusAuthorized {
		return nil, fmt.Errorf("payment cannot be captured in status: %s", payment.Status)
	}

	result, err := s.gateway.Capture(ctx, payment.GatewayID, payment.Amount)
	if err != nil {
		_ = s.repo.UpdateStatus(ctx, payment.ID, model.PaymentStatusFailed, payment.GatewayID, "", err.Error())
		payment.Status = model.PaymentStatusFailed
		payment.FailureReason = err.Error()
		_ = s.publisher.PublishPaymentFailed(ctx, payment)
		return payment, fmt.Errorf("gateway capture failed: %w", err)
	}

	_ = s.repo.UpdateStatus(ctx, payment.ID, model.PaymentStatusCaptured, payment.GatewayID, result.GatewayStatus, "")
	payment.Status = model.PaymentStatusCaptured
	payment.GatewayStatus = result.GatewayStatus

	_ = s.publisher.PublishPaymentCaptured(ctx, payment)
	return payment, nil
}

func (s *paymentService) Refund(ctx context.Context, req dto.RefundRequest) (*model.Refund, error) {
	payment, err := s.repo.FindByID(ctx, req.PaymentID)
	if err != nil {
		return nil, err
	}

	if payment.Status != model.PaymentStatusCaptured {
		return nil, fmt.Errorf("payment cannot be refunded in status: %s", payment.Status)
	}

	_, err = s.gateway.Refund(ctx, payment.GatewayID, req.Amount)
	if err != nil {
		return nil, fmt.Errorf("gateway refund failed: %w", err)
	}

	refund := &model.Refund{
		ID:        uuid.New().String(),
		PaymentID: req.PaymentID,
		Amount:    req.Amount,
		Reason:    req.Reason,
		Status:    model.RefundStatusProcessed,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.repo.CreateRefund(ctx, refund); err != nil {
		return nil, err
	}

	_ = s.repo.UpdateStatus(ctx, payment.ID, model.PaymentStatusRefunded, payment.GatewayID, payment.GatewayStatus, "")
	_ = s.publisher.PublishRefundIssued(ctx, refund)
	return refund, nil
}

func (s *paymentService) GetStatus(ctx context.Context, paymentID string) (*model.Payment, error) {
	return s.repo.FindByID(ctx, paymentID)
}
