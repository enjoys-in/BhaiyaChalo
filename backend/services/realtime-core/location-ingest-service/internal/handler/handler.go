package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/location-ingest-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/location-ingest-service/internal/ports"
)

type LocationIngestHandler struct {
	svc ports.LocationIngestService
}

func NewLocationIngestHandler(svc ports.LocationIngestService) *LocationIngestHandler {
	return &LocationIngestHandler{svc: svc}
}

func (h *LocationIngestHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateLocationUpdateRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusCreated, "locationIngest created successfully", resp)
}

func (h *LocationIngestHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	successJSON(w, http.StatusOK, "locationIngest fetched successfully", resp)
}

func (h *LocationIngestHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateLocationUpdateRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "locationIngest updated successfully", resp)
}

func (h *LocationIngestHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "locationIngest deleted successfully", nil)
}
