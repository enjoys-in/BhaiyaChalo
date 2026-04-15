package handler

import (
	"encoding/json"
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/eta-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/eta-service/internal/ports"
)

type Handler struct {
	svc ports.ETAService
}

func NewHandler(svc ports.ETAService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CalculateETA(w http.ResponseWriter, r *http.Request) {
	var req dto.CalculateETARequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := validateCalculateETARequest(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.svc.Calculate(r.Context(), &req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := &dto.ETAResponse{
		DistanceKM:        result.DistanceKM,
		DurationMinutes:   result.DurationMinutes,
		TrafficMultiplier: result.TrafficMultiplier,
		CalculatedAt:      result.CalculatedAt,
		VehicleType:       req.VehicleType,
		CityID:            req.CityID,
	}

	writeJSON(w, http.StatusOK, resp)
}
