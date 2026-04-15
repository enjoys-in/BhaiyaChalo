package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *OTPHandler) {
	mux.HandleFunc("POST /api/v1/otp/send", h.SendOTP)
	mux.HandleFunc("POST /api/v1/otp/verify", h.VerifyOTP)
}
