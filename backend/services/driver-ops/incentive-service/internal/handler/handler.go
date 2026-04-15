package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/incentive-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/incentive-service/internal/ports"
)

type IncentiveHandler struct {
	svc ports.IncentiveService
}

func NewIncentiveHandler(svc ports.IncentiveService) *IncentiveHandler {
	return &IncentiveHandler{svc: svc}
}

func (h *IncentiveHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateIncentiveRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusCreated, "incentive created successfully", resp)
}

func (h *IncentiveHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	successJSON(w, http.StatusOK, "incentive fetched successfully", resp)
}

func (h *IncentiveHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateIncentiveRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "incentive updated successfully", resp)
}

func (h *IncentiveHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "incentive deleted successfully", nil)
}
