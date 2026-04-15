package handler

import (
	"encoding/json"
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/ports"
)

type Handler struct {
	svc ports.TrackingService
}

func NewHandler(svc ports.TrackingService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := validateUpdateLocationRequest(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.svc.UpdateLocation(r.Context(), &req); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) GetLocation(w http.ResponseWriter, r *http.Request) {
	driverID := extractPathParam(r, "driverID")
	if driverID == "" {
		writeError(w, http.StatusBadRequest, "driver_id is required")
		return
	}

	loc, err := h.svc.GetLocation(r.Context(), driverID)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, toLocationResponse(loc))
}

func (h *Handler) StartTracking(w http.ResponseWriter, r *http.Request) {
	var req dto.StartTrackingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := validateStartTrackingRequest(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	session, err := h.svc.StartTracking(r.Context(), &req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, toTrackingSessionResponse(session))
}

func (h *Handler) StopTracking(w http.ResponseWriter, r *http.Request) {
	tripID := extractPathParam(r, "tripID")
	if tripID == "" {
		writeError(w, http.StatusBadRequest, "trip_id is required")
		return
	}

	if err := h.svc.StopTracking(r.Context(), tripID); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "stopped"})
}
