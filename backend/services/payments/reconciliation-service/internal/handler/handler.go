package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/ports"
)

type ReconciliationHandler struct {
	svc ports.ReconciliationService
}

func NewReconciliationHandler(svc ports.ReconciliationService) *ReconciliationHandler {
	return &ReconciliationHandler{svc: svc}
}

func (h *ReconciliationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateReconciliationRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusCreated, "reconciliation created successfully", resp)
}

func (h *ReconciliationHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	successJSON(w, http.StatusOK, "reconciliation fetched successfully", resp)
}

func (h *ReconciliationHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateReconciliationRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "reconciliation updated successfully", resp)
}

func (h *ReconciliationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "reconciliation deleted successfully", nil)
}
