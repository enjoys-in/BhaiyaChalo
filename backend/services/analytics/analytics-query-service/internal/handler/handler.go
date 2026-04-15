package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-query-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-query-service/internal/ports"
)

type AnalyticsQueryHandler struct {
	svc ports.AnalyticsQueryService
}

func NewAnalyticsQueryHandler(svc ports.AnalyticsQueryService) *AnalyticsQueryHandler {
	return &AnalyticsQueryHandler{svc: svc}
}

func (h *AnalyticsQueryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateQueryResultRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusCreated, "analyticsQuery created successfully", resp)
}

func (h *AnalyticsQueryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	successJSON(w, http.StatusOK, "analyticsQuery fetched successfully", resp)
}

func (h *AnalyticsQueryHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateQueryResultRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "analyticsQuery updated successfully", resp)
}

func (h *AnalyticsQueryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "analyticsQuery deleted successfully", nil)
}
