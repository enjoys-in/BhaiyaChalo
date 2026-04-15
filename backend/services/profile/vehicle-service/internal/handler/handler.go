package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/vehicle-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/vehicle-service/internal/ports"
)

type VehicleHandler struct {
	svc ports.VehicleService
}

func NewVehicleHandler(svc ports.VehicleService) *VehicleHandler {
	return &VehicleHandler{svc: svc}
}

func (h *VehicleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateVehicleRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.DriverID == "" || req.Make == "" || req.PlateNumber == "" {
		errorJSON(w, http.StatusBadRequest, "driver_id, make and plate_number are required")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusCreated, "", resp)
}

func (h *VehicleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "vehicle id is required")
		return
	}

	resp, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		errorJSON(w, http.StatusNotFound, "vehicle not found")
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *VehicleHandler) GetByDriverID(w http.ResponseWriter, r *http.Request) {
	driverID := r.PathValue("driverId")
	if driverID == "" {
		errorJSON(w, http.StatusBadRequest, "driver id is required")
		return
	}

	resp, err := h.svc.GetByDriverID(r.Context(), driverID)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *VehicleHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "vehicle id is required")
		return
	}

	var req dto.UpdateVehicleRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), id, req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *VehicleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "vehicle id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", map[string]string{"status": "deleted"})
}

func (h *VehicleHandler) ListByType(w http.ResponseWriter, r *http.Request) {
	vehicleType := r.URL.Query().Get("type")
	if vehicleType == "" {
		errorJSON(w, http.StatusBadRequest, "type query parameter is required")
		return
	}

	resp, err := h.svc.ListByType(r.Context(), vehicleType)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *VehicleHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	successJSON(w, http.StatusOK, "", map[string]string{"status": "ok"})
}
