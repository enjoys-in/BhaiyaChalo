package handler

import (
	"encoding/json"
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/stop-planning-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/stop-planning-service/internal/ports"
)

type Handler struct {
	svc ports.StopService
}

func NewHandler(svc ports.StopService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) AddStop(w http.ResponseWriter, r *http.Request) {
	var req dto.AddStopRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.TripID == "" {
		writeError(w, http.StatusBadRequest, "trip_id is required")
		return
	}

	stop, err := h.svc.AddStop(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, toStopResponse(stop))
}

func (h *Handler) RemoveStop(w http.ResponseWriter, r *http.Request) {
	tripID := extractPathParam(r, "tripId")
	stopID := extractPathParam(r, "stopId")

	req := dto.RemoveStopRequest{TripID: tripID, StopID: stopID}
	if err := h.svc.RemoveStop(r.Context(), req); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) ReorderStops(w http.ResponseWriter, r *http.Request) {
	var req dto.ReorderStopsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	tripID := extractPathParam(r, "tripId")
	req.TripID = tripID

	result, err := h.svc.ReorderStops(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, toMultiStopResponse(result))
}

func (h *Handler) UpdateStopStatus(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateStopStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	req.TripID = extractPathParam(r, "tripId")
	req.StopID = extractPathParam(r, "stopId")

	stop, err := h.svc.UpdateStopStatus(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, toStopResponse(stop))
}

func (h *Handler) GetStopsByTrip(w http.ResponseWriter, r *http.Request) {
	tripID := extractPathParam(r, "tripId")
	if tripID == "" {
		writeError(w, http.StatusBadRequest, "trip_id is required")
		return
	}

	result, err := h.svc.GetStopsByTrip(r.Context(), tripID)
	if err != nil {
		writeError(w, http.StatusNotFound, "stops not found")
		return
	}

	writeJSON(w, http.StatusOK, toMultiStopResponse(result))
}
