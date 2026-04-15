package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/trip-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/trip-service/internal/ports"
)

type Handler struct {
	svc  ports.TripService
	repo ports.TripRepository
}

func NewHandler(svc ports.TripService, repo ports.TripRepository) *Handler {
	return &Handler{svc: svc, repo: repo}
}

func (h *Handler) CreateTrip(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTripRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.BookingID == "" || req.UserID == "" || req.DriverID == "" {
		writeError(w, http.StatusBadRequest, "booking_id, user_id and driver_id are required")
		return
	}

	trip, err := h.svc.Create(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, toTripResponse(trip))
}

func (h *Handler) GetTrip(w http.ResponseWriter, r *http.Request) {
	tripID := extractPathParam(r, "tripId")
	if tripID == "" {
		writeError(w, http.StatusBadRequest, "trip_id is required")
		return
	}

	trip, err := h.svc.Get(r.Context(), tripID)
	if err != nil {
		writeError(w, http.StatusNotFound, "trip not found")
		return
	}

	writeJSON(w, http.StatusOK, toTripResponse(trip))
}

func (h *Handler) UpdateTripStatus(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateTripStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	req.TripID = extractPathParam(r, "tripId")

	if err := h.svc.UpdateStatus(r.Context(), req); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) ListByUser(w http.ResponseWriter, r *http.Request) {
	userID := extractPathParam(r, "userId")
	limit, offset := parsePagination(r)

	trips, err := h.svc.ListByUser(r.Context(), userID, limit, offset)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var resp []interface{}
	for _, t := range trips {
		resp = append(resp, toTripResponse(&t))
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) ListByDriver(w http.ResponseWriter, r *http.Request) {
	driverID := extractPathParam(r, "driverId")
	limit, offset := parsePagination(r)

	trips, err := h.svc.ListByDriver(r.Context(), driverID, limit, offset)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var resp []interface{}
	for _, t := range trips {
		resp = append(resp, toTripResponse(&t))
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) GetTimeline(w http.ResponseWriter, r *http.Request) {
	tripID := extractPathParam(r, "tripId")

	events, err := h.repo.GetTimeline(r.Context(), tripID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := make([]dto.TripTimelineResponse, len(events))
	for i, e := range events {
		resp[i] = dto.TripTimelineResponse{
			TripID:    e.TripID,
			Event:     e.Event,
			Timestamp: e.Timestamp,
		}
	}
	writeJSON(w, http.StatusOK, resp)
}

func parsePagination(r *http.Request) (int, int) {
	limit := 20
	offset := 0
	if v := r.URL.Query().Get("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 && n <= 100 {
			limit = n
		}
	}
	if v := r.URL.Query().Get("offset"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 0 {
			offset = n
		}
	}
	return limit, offset
}
