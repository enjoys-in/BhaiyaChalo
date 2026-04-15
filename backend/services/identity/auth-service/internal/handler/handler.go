package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/auth-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/auth-service/internal/ports"
)

type AuthHandler struct {
	svc ports.AuthService
}

func NewAuthHandler(svc ports.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Login(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "login successful", resp)
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var req dto.RefreshRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Refresh(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusUnauthorized, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "token refreshed", resp)
}

func (h *AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {
	var req dto.VerifyRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Verify(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusUnauthorized, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "token valid", resp)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var req dto.LogoutRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.svc.Logout(r.Context(), req); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "logged out successfully", nil)
}
