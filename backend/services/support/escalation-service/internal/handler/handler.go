package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/ports"
)

type EscalationHandler struct {
	svc ports.EscalationService
}

func NewEscalationHandler(svc ports.EscalationService) *EscalationHandler {
	return &EscalationHandler{svc: svc}
}

func (h *EscalationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateEscalationRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusCreated, "escalation created successfully", resp)
}

func (h *EscalationHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	successJSON(w, http.StatusOK, "escalation fetched successfully", resp)
}

func (h *EscalationHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateEscalationRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "escalation updated successfully", resp)
}

func (h *EscalationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "escalation deleted successfully", nil)
}
