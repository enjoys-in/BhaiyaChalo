package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/referral-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/referral-service/internal/ports"
)

type ReferralHandler struct {
	svc ports.ReferralService
}

func NewReferralHandler(svc ports.ReferralService) *ReferralHandler {
	return &ReferralHandler{svc: svc}
}

func (h *ReferralHandler) GenerateCode(w http.ResponseWriter, r *http.Request) {
	var req dto.GenerateCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.UserID == "" {
		writeError(w, http.StatusBadRequest, "user_id is required")
		return
	}

	resp, err := h.svc.GenerateCode(r.Context(), &req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, resp)
}

func (h *ReferralHandler) ApplyReferral(w http.ResponseWriter, r *http.Request) {
	var req dto.ApplyReferralRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := validateApplyReferralRequest(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.svc.ApplyReferral(r.Context(), &req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *ReferralHandler) CompleteReferral(w http.ResponseWriter, r *http.Request) {
	referralID := extractPathParam(r.URL.Path, "/api/v1/referrals/", "/complete")
	if referralID == "" {
		writeError(w, http.StatusBadRequest, "referral_id is required")
		return
	}

	resp, err := h.svc.CompleteReferral(r.Context(), referralID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *ReferralHandler) ClaimReward(w http.ResponseWriter, r *http.Request) {
	var req dto.ClaimRewardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.ReferralID == "" || req.UserID == "" {
		writeError(w, http.StatusBadRequest, "referral_id and user_id are required")
		return
	}

	resp, err := h.svc.ClaimReward(r.Context(), &req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *ReferralHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		writeError(w, http.StatusBadRequest, "user_id query parameter is required")
		return
	}

	resp, err := h.svc.GetStats(r.Context(), userID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *ReferralHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func extractPathParam(path, prefix, suffix string) string {
	path = strings.TrimPrefix(path, prefix)
	path = strings.TrimSuffix(path, suffix)
	return path
}
