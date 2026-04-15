package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *SurgeHandler) {
	mux.HandleFunc("POST /api/v1/surge/calculate", h.Calculate)
	mux.HandleFunc("GET /api/v1/surge/zones/{zoneId}", h.GetCurrentSurge)
	mux.HandleFunc("PUT /api/v1/surge/policies/{cityId}", h.UpdatePolicy)
}
