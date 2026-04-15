package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/realtime-metrics-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/analytics/realtime-metrics-service/internal/ports"
)

type RealtimeMetricsHandler struct {
	svc ports.RealtimeMetricsService
}

func NewRealtimeMetricsHandler(svc ports.RealtimeMetricsService) *RealtimeMetricsHandler {
	return &RealtimeMetricsHandler{svc: svc}
}

func (h *RealtimeMetricsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateMetricRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusCreated, "realtimeMetrics created successfully", resp)
}

func (h *RealtimeMetricsHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	successJSON(w, http.StatusOK, "realtimeMetrics fetched successfully", resp)
}

func (h *RealtimeMetricsHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateMetricRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "realtimeMetrics updated successfully", resp)
}

func (h *RealtimeMetricsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "realtimeMetrics deleted successfully", nil)
}
