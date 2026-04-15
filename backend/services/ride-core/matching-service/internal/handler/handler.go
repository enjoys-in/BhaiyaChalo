package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/matching-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/matching-service/internal/ports"
)

type MatchingHandler struct {
	svc ports.MatchingService
}

func NewMatchingHandler(svc ports.MatchingService) *MatchingHandler {
	return &MatchingHandler{svc: svc}
}

func (h *MatchingHandler) FindDrivers(w http.ResponseWriter, r *http.Request) {
	var req dto.FindDriversRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.BookingID == "" || req.CityID == "" || req.VehicleType == "" {
		errorJSON(w, http.StatusBadRequest, "booking_id, city_id, and vehicle_type are required")
		return
	}

	resp, err := h.svc.FindNearestDrivers(r.Context(), &req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *MatchingHandler) AssignDriver(w http.ResponseWriter, r *http.Request) {
	var req dto.FindDriversRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.BookingID == "" || req.CityID == "" || req.VehicleType == "" {
		errorJSON(w, http.StatusBadRequest, "booking_id, city_id, and vehicle_type are required")
		return
	}

	result, err := h.svc.AssignBestDriver(r.Context(), &req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "driver matched successfully", dto.MatchResponse{
		DriverID: result.DriverID,
		Distance: result.DistanceKM,
		ETA:      result.ETASeconds,
	})
}
