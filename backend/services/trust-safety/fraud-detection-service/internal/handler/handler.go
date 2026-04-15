package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/fraud-detection-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/fraud-detection-service/internal/ports"
)

type FraudDetectionHandler struct {
	svc ports.FraudDetectionService
}

func NewFraudDetectionHandler(svc ports.FraudDetectionService) *FraudDetectionHandler {
	return &FraudDetectionHandler{svc: svc}
}

func (h *FraudDetectionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateFraudSignalRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusCreated, "fraudDetection created successfully", resp)
}

func (h *FraudDetectionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	successJSON(w, http.StatusOK, "fraudDetection fetched successfully", resp)
}

func (h *FraudDetectionHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateFraudSignalRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "fraudDetection updated successfully", resp)
}

func (h *FraudDetectionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "fraudDetection deleted successfully", nil)
}
