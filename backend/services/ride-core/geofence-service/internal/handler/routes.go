package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *GeofenceHandler) {
	mux.HandleFunc("POST /api/v1/geofences", h.Create)
	mux.HandleFunc("GET /api/v1/geofences", h.List)
	mux.HandleFunc("GET /api/v1/geofences/{id}", h.Get)
	mux.HandleFunc("PUT /api/v1/geofences/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/geofences/{id}", h.Delete)
	mux.HandleFunc("POST /api/v1/geofences/check-point", h.CheckPoint)
}
