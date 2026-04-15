package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-ingestion-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-ingestion-service/internal/ports"
)

type AnalyticsIngestionHandler struct {
	svc ports.AnalyticsIngestionService
}

func NewAnalyticsIngestionHandler(svc ports.AnalyticsIngestionService) *AnalyticsIngestionHandler {
	return &AnalyticsIngestionHandler{svc: svc}
}

func (h *AnalyticsIngestionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateAnalyticsEventRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusCreated, "analyticsIngestion created successfully", resp)
}

func (h *AnalyticsIngestionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	successJSON(w, http.StatusOK, "analyticsIngestion fetched successfully", resp)
}

func (h *AnalyticsIngestionHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateAnalyticsEventRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "analyticsIngestion updated successfully", resp)
}

func (h *AnalyticsIngestionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "analyticsIngestion deleted successfully", nil)
}
