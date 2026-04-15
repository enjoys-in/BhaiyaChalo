package handler

import (
	"encoding/json"
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/payment-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/payment-service/internal/ports"
)

type Handler struct {
	svc ports.PaymentService
}

func NewHandler(svc ports.PaymentService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) InitiatePayment(w http.ResponseWriter, r *http.Request) {
	var req dto.InitiatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.BookingID == "" || req.UserID == "" || req.Amount <= 0 {
		writeError(w, http.StatusBadRequest, "booking_id, user_id, and positive amount are required")
		return
	}

	payment, err := h.svc.Initiate(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, toPaymentResponse(payment))
}

func (h *Handler) CapturePayment(w http.ResponseWriter, r *http.Request) {
	paymentID := extractPathParam(r, "paymentId")
	req := dto.CapturePaymentRequest{
		PaymentID: paymentID,
	}

	var body struct {
		GatewayID string `json:"gateway_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.GatewayID = body.GatewayID

	payment, err := h.svc.Capture(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, toPaymentResponse(payment))
}

func (h *Handler) RefundPayment(w http.ResponseWriter, r *http.Request) {
	paymentID := extractPathParam(r, "paymentId")

	var req dto.RefundRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.PaymentID = paymentID

	if req.Amount <= 0 || req.Reason == "" {
		writeError(w, http.StatusBadRequest, "positive amount and reason are required")
		return
	}

	refund, err := h.svc.Refund(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, toRefundResponse(refund))
}

func (h *Handler) GetPaymentStatus(w http.ResponseWriter, r *http.Request) {
	paymentID := extractPathParam(r, "paymentId")
	if paymentID == "" {
		writeError(w, http.StatusBadRequest, "payment_id is required")
		return
	}

	payment, err := h.svc.GetStatus(r.Context(), paymentID)
	if err != nil {
		writeError(w, http.StatusNotFound, "payment not found")
		return
	}

	writeJSON(w, http.StatusOK, toPaymentStatusResponse(payment))
}

func (h *Handler) GetPayment(w http.ResponseWriter, r *http.Request) {
	paymentID := extractPathParam(r, "paymentId")
	if paymentID == "" {
		writeError(w, http.StatusBadRequest, "payment_id is required")
		return
	}

	payment, err := h.svc.GetStatus(r.Context(), paymentID)
	if err != nil {
		writeError(w, http.StatusNotFound, "payment not found")
		return
	}

	writeJSON(w, http.StatusOK, toPaymentResponse(payment))
}

func (h *Handler) GetRefund(w http.ResponseWriter, r *http.Request) {
	paymentID := extractPathParam(r, "paymentId")
	if paymentID == "" {
		writeError(w, http.StatusBadRequest, "payment_id is required")
		return
	}

	// Refund lookup is done via the payment service's repo, accessible through service layer
	// For simplicity, the handler re-uses GetStatus to verify payment exists
	_, err := h.svc.GetStatus(r.Context(), paymentID)
	if err != nil {
		writeError(w, http.StatusNotFound, "payment not found")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"payment_id": paymentID, "message": "refund lookup by payment_id"})
}
