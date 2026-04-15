package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/pricing-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/pricing-service/internal/ports"
)

type PricingHandler struct {
	svc ports.PricingService
}

func NewPricingHandler(svc ports.PricingService) *PricingHandler {
	return &PricingHandler{svc: svc}
}

// Estimate handles POST /api/v1/pricing/estimate — gRPC-ready, called by booking-service.
func (h *PricingHandler) Estimate(w http.ResponseWriter, r *http.Request) {
	var req dto.EstimatePriceRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.CityID == "" || req.VehicleType == "" {
		errorJSON(w, http.StatusBadRequest, "city_id and vehicle_type are required")
		return
	}
	if req.DistanceKM <= 0 || req.DurationMin <= 0 {
		errorJSON(w, http.StatusBadRequest, "distance_km and duration_min must be positive")
		return
	}

	resp, err := h.svc.Estimate(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

// GetRules handles GET /api/v1/pricing/rules?city_id=...
func (h *PricingHandler) GetRules(w http.ResponseWriter, r *http.Request) {
	cityID := r.URL.Query().Get("city_id")
	if cityID == "" {
		errorJSON(w, http.StatusBadRequest, "city_id query parameter is required")
		return
	}

	resp, err := h.svc.GetRules(r.Context(), cityID)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

// CreateRule handles POST /api/v1/pricing/rules
func (h *PricingHandler) CreateRule(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePricingRuleRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.CityID == "" || req.VehicleType == "" {
		errorJSON(w, http.StatusBadRequest, "city_id and vehicle_type are required")
		return
	}
	if req.BaseFarePerKM <= 0 || req.BaseFarePerMin <= 0 {
		errorJSON(w, http.StatusBadRequest, "fare rates must be positive")
		return
	}

	resp, err := h.svc.CreateRule(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusCreated, "", resp)
}

// UpdateRule handles PUT /api/v1/pricing/rules/{id}
func (h *PricingHandler) UpdateRule(w http.ResponseWriter, r *http.Request) {
	id := pathParam(r, "id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "rule id is required")
		return
	}

	var req dto.UpdatePricingRuleRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.UpdateRule(r.Context(), id, req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}
