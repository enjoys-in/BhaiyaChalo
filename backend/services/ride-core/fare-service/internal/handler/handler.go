package handler

import (
	"encoding/json"
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/fare-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/fare-service/internal/ports"
)

type Handler struct {
	svc ports.FareService
}

func NewHandler(svc ports.FareService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CalculateFare(w http.ResponseWriter, r *http.Request) {
	var req dto.CalculateFareRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.BookingID == "" || req.CityID == "" || req.VehicleType == "" {
		writeError(w, http.StatusBadRequest, "booking_id, city_id and vehicle_type are required")
		return
	}

	if req.DistanceKM <= 0 || req.DurationMin <= 0 {
		writeError(w, http.StatusBadRequest, "distance_km and duration_min must be positive")
		return
	}

	calc, err := h.svc.Calculate(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, toFareBreakdownResponse(calc))
}

func (h *Handler) RecalculateFare(w http.ResponseWriter, r *http.Request) {
	var req dto.RecalculateFareRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.BookingID == "" {
		writeError(w, http.StatusBadRequest, "booking_id is required")
		return
	}

	calc, err := h.svc.Recalculate(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, toFareBreakdownResponse(calc))
}

func (h *Handler) GetBreakdown(w http.ResponseWriter, r *http.Request) {
	bookingID := extractPathParam(r, "bookingId")
	if bookingID == "" {
		writeError(w, http.StatusBadRequest, "booking_id is required")
		return
	}

	calc, err := h.svc.GetBreakdown(r.Context(), bookingID)
	if err != nil {
		writeError(w, http.StatusNotFound, "fare breakdown not found")
		return
	}

	writeJSON(w, http.StatusOK, toFareBreakdownResponse(calc))
}
