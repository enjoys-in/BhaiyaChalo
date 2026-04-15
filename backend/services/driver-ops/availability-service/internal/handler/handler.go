package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/ports"
)

type AvailabilityHandler struct {
	svc ports.AvailabilityService
}

func NewAvailabilityHandler(svc ports.AvailabilityService) *AvailabilityHandler {
	return &AvailabilityHandler{svc: svc}
}

// GoOnline handles POST /api/v1/availability/online
func (h *AvailabilityHandler) GoOnline(w http.ResponseWriter, r *http.Request) {
	var req dto.GoOnlineRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.DriverID == "" || req.CityID == "" || req.VehicleType == "" {
		errorJSON(w, http.StatusBadRequest, "driver_id, city_id and vehicle_type are required")
		return
	}

	resp, err := h.svc.GoOnline(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

// GoOffline handles POST /api/v1/availability/offline
func (h *AvailabilityHandler) GoOffline(w http.ResponseWriter, r *http.Request) {
	var req dto.GoOfflineRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.DriverID == "" {
		errorJSON(w, http.StatusBadRequest, "driver_id is required")
		return
	}

	if err := h.svc.GoOffline(r.Context(), req); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", map[string]string{"status": "offline"})
}

// UpdateTripStatus handles PUT /api/v1/availability/trip-status
func (h *AvailabilityHandler) UpdateTripStatus(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateTripStatusRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.DriverID == "" {
		errorJSON(w, http.StatusBadRequest, "driver_id is required")
		return
	}

	if err := h.svc.UpdateTripStatus(r.Context(), req); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", map[string]string{"status": "updated"})
}

// GetStatus handles GET /api/v1/availability/{driverID}
func (h *AvailabilityHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	driverID := pathParam(r, "driverID")
	if driverID == "" {
		errorJSON(w, http.StatusBadRequest, "driver_id is required")
		return
	}

	resp, err := h.svc.GetStatus(r.Context(), driverID)
	if err != nil {
		errorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

// CountOnline handles GET /api/v1/availability/count?city_id=...&vehicle_type=...
func (h *AvailabilityHandler) CountOnline(w http.ResponseWriter, r *http.Request) {
	cityID := r.URL.Query().Get("city_id")
	vehicleType := r.URL.Query().Get("vehicle_type")

	if cityID == "" {
		errorJSON(w, http.StatusBadRequest, "city_id query parameter is required")
		return
	}

	resp, err := h.svc.CountOnlineDrivers(r.Context(), cityID, vehicleType)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}
