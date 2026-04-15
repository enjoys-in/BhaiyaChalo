package handler

import (
	"encoding/json"
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/route-planning-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/route-planning-service/internal/ports"
)

type Handler struct {
	svc ports.RouteService
}

func NewHandler(svc ports.RouteService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) PlanRoute(w http.ResponseWriter, r *http.Request) {
	var req dto.PlanRouteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.BookingID == "" {
		writeError(w, http.StatusBadRequest, "booking_id is required")
		return
	}

	route, err := h.svc.PlanRoute(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := toRouteResponse(route)
	writeJSON(w, http.StatusCreated, resp)
}

func (h *Handler) GetRoute(w http.ResponseWriter, r *http.Request) {
	bookingID := extractPathParam(r, "bookingId")
	if bookingID == "" {
		writeError(w, http.StatusBadRequest, "booking_id is required")
		return
	}

	route, err := h.svc.GetRoute(r.Context(), bookingID)
	if err != nil {
		writeError(w, http.StatusNotFound, "route not found")
		return
	}

	resp := toRouteResponse(route)
	writeJSON(w, http.StatusOK, resp)
}
