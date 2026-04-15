package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/risk-scoring-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/risk-scoring-service/internal/ports"
)

type RiskScoringHandler struct {
	svc ports.RiskScoringService
}

func NewRiskScoringHandler(svc ports.RiskScoringService) *RiskScoringHandler {
	return &RiskScoringHandler{svc: svc}
}

func (h *RiskScoringHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateRiskScoreRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusCreated, "riskScoring created successfully", resp)
}

func (h *RiskScoringHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	resp, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "riskScoring fetched successfully", resp)
}

func (h *RiskScoringHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateRiskScoreRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "riskScoring updated successfully", resp)
}

func (h *RiskScoringHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "riskScoring deleted successfully", nil)
}
