package handler

import (
	"errors"
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/ports"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/service"
)

type OTPHandler struct {
	svc ports.OTPService
}

func NewOTPHandler(svc ports.OTPService) *OTPHandler {
	return &OTPHandler{svc: svc}
}

func (h *OTPHandler) SendOTP(w http.ResponseWriter, r *http.Request) {
	var req dto.SendOTPRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Phone == "" || req.Purpose == "" {
		errorJSON(w, http.StatusBadRequest, "phone and purpose are required")
		return
	}

	resp, err := h.svc.Send(r.Context(), req)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *OTPHandler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req dto.VerifyOTPRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Phone == "" || req.Code == "" || req.Purpose == "" {
		errorJSON(w, http.StatusBadRequest, "phone, code and purpose are required")
		return
	}

	resp, err := h.svc.Verify(r.Context(), req)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func handleServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrRateLimited):
		errorJSON(w, http.StatusTooManyRequests, err.Error())
	case errors.Is(err, service.ErrOTPExpired):
		errorJSON(w, http.StatusGone, err.Error())
	case errors.Is(err, service.ErrOTPInvalid):
		errorJSON(w, http.StatusUnprocessableEntity, err.Error())
	case errors.Is(err, service.ErrMaxAttemptsExceed):
		errorJSON(w, http.StatusForbidden, err.Error())
	case errors.Is(err, service.ErrOTPNotFound):
		errorJSON(w, http.StatusNotFound, err.Error())
	default:
		errorJSON(w, http.StatusInternalServerError, "internal server error")
	}
}
