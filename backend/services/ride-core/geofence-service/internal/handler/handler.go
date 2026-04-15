package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/geofence-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/geofence-service/internal/ports"
)

type GeofenceHandler struct {
	svc ports.GeofenceService
}

func NewGeofenceHandler(svc ports.GeofenceService) *GeofenceHandler {
	return &GeofenceHandler{svc: svc}
}

// Create handles POST /api/v1/geofences
func (h *GeofenceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateGeofenceRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.CityID == "" || req.Name == "" || req.Type == "" {
		errorJSON(w, http.StatusBadRequest, "city_id, name and type are required")
		return
	}
	if len(req.Polygon) < 3 {
		errorJSON(w, http.StatusBadRequest, "polygon must have at least 3 coordinates")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusCreated, "", resp)
}

// Get handles GET /api/v1/geofences/{id}
func (h *GeofenceHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := pathParam(r, "id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "geofence id is required")
		return
	}

	resp, err := h.svc.Get(r.Context(), id)
	if err != nil {
		errorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

// List handles GET /api/v1/geofences?city_id=...
func (h *GeofenceHandler) List(w http.ResponseWriter, r *http.Request) {
	cityID := r.URL.Query().Get("city_id")
	if cityID == "" {
		errorJSON(w, http.StatusBadRequest, "city_id query parameter is required")
		return
	}

	resp, err := h.svc.List(r.Context(), cityID)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

// Update handles PUT /api/v1/geofences/{id}
func (h *GeofenceHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := pathParam(r, "id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "geofence id is required")
		return
	}

	var req dto.UpdateGeofenceRequest
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

// Delete handles DELETE /api/v1/geofences/{id}
func (h *GeofenceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := pathParam(r, "id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "geofence id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", map[string]string{"status": "deleted"})
}

// CheckPoint handles POST /api/v1/geofences/check-point
func (h *GeofenceHandler) CheckPoint(w http.ResponseWriter, r *http.Request) {
	var req dto.PointInFenceRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Lat == 0 && req.Lng == 0 {
		errorJSON(w, http.StatusBadRequest, "lat and lng are required")
		return
	}

	resp, err := h.svc.CheckPoint(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}
