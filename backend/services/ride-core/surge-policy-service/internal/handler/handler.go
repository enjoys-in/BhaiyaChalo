package handler

import (
	"net/http"
	"strconv"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/surge-policy-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/surge-policy-service/internal/ports"
)

type SurgeHandler struct {
	svc ports.SurgeService
}

func NewSurgeHandler(svc ports.SurgeService) *SurgeHandler {
	return &SurgeHandler{svc: svc}
}

// Calculate handles POST /api/v1/surge/calculate
func (h *SurgeHandler) Calculate(w http.ResponseWriter, r *http.Request) {
	var req dto.CalculateSurgeRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.CityID == "" || req.ZoneID == "" {
		errorJSON(w, http.StatusBadRequest, "city_id and zone_id are required")
		return
	}
	if req.DemandCount < 0 || req.SupplyCount < 0 {
		errorJSON(w, http.StatusBadRequest, "demand_count and supply_count must be non-negative")
		return
	}

	resp, err := h.svc.Calculate(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

// GetCurrentSurge handles GET /api/v1/surge/zones/{zoneId}
func (h *SurgeHandler) GetCurrentSurge(w http.ResponseWriter, r *http.Request) {
	zoneID := pathParam(r, "zoneId")
	if zoneID == "" {
		errorJSON(w, http.StatusBadRequest, "zone id is required")
		return
	}

	resp, err := h.svc.GetCurrentSurge(r.Context(), zoneID)
	if err != nil {
		errorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

// UpdatePolicy handles PUT /api/v1/surge/policies/{cityId}
func (h *SurgeHandler) UpdatePolicy(w http.ResponseWriter, r *http.Request) {
	cityID := pathParam(r, "cityId")
	if cityID == "" {
		errorJSON(w, http.StatusBadRequest, "city id is required")
		return
	}

	var req dto.UpdatePolicyRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.UpdatePolicy(r.Context(), cityID, req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

// parseIntParam parses a query string integer parameter with a default value.
func parseIntParam(r *http.Request, key string, defaultVal int) int {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultVal
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return n
}
