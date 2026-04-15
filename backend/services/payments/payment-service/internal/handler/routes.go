package handler

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/payments", h.InitiatePayment)
	mux.HandleFunc("GET /api/v1/payments/{paymentId}", h.GetPayment)
	mux.HandleFunc("GET /api/v1/payments/{paymentId}/status", h.GetPaymentStatus)
	mux.HandleFunc("POST /api/v1/payments/{paymentId}/capture", h.CapturePayment)
	mux.HandleFunc("POST /api/v1/payments/{paymentId}/refund", h.RefundPayment)
}
