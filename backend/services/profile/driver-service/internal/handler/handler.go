package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/driver-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/driver-service/internal/ports"
)

type DriverHandler struct {
	svc ports.DriverService
}

func NewDriverHandler(svc ports.DriverService) *DriverHandler {
	return &DriverHandler{svc: svc}
}

func (h *DriverHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateDriverRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Phone == "" || req.FirstName == "" {
		errorJSON(w, http.StatusBadRequest, "phone and first_name are required")
		return
	}

	resp, err := h.svc.Create(r.Context(), &req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "failed to create driver")
		return
	}

	successJSON(w, http.StatusCreated, "", resp)
}

func (h *DriverHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "driver id is required")
		return
	}

	resp, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		errorJSON(w, http.StatusNotFound, "driver not found")
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *DriverHandler) GetByPhone(w http.ResponseWriter, r *http.Request) {
	phone := r.PathValue("phone")
	if phone == "" {
		errorJSON(w, http.StatusBadRequest, "phone is required")
		return
	}

	resp, err := h.svc.GetByPhone(r.Context(), phone)
	if err != nil {
		errorJSON(w, http.StatusNotFound, "driver not found")
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *DriverHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "driver id is required")
		return
	}

	var req dto.UpdateDriverRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), id, &req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "failed to update driver")
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *DriverHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "driver id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, "failed to delete driver")
		return
	}

	successJSON(w, http.StatusOK, "", map[string]string{"status": "deleted"})
}

func (h *DriverHandler) ListByCityID(w http.ResponseWriter, r *http.Request) {
	cityID := r.PathValue("cityId")
	if cityID == "" {
		errorJSON(w, http.StatusBadRequest, "city id is required")
		return
	}

	resp, err := h.svc.ListByCityID(r.Context(), cityID)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "failed to list drivers")
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *DriverHandler) GetPreference(w http.ResponseWriter, r *http.Request) {
	driverID := r.PathValue("id")
	if driverID == "" {
		errorJSON(w, http.StatusBadRequest, "driver id is required")
		return
	}

	resp, err := h.svc.GetPreference(r.Context(), driverID)
	if err != nil {
		errorJSON(w, http.StatusNotFound, "preferences not found")
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *DriverHandler) UpdatePreference(w http.ResponseWriter, r *http.Request) {
	driverID := r.PathValue("id")
	if driverID == "" {
		errorJSON(w, http.StatusBadRequest, "driver id is required")
		return
	}

	var req dto.UpdatePreferenceRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.UpdatePreference(r.Context(), driverID, &req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "failed to update preferences")
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}
