package handler

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/tracking/location", h.UpdateLocation)
	mux.HandleFunc("GET /api/v1/tracking/location/{driverID}", h.GetLocation)
	mux.HandleFunc("POST /api/v1/tracking/session/start", h.StartTracking)
	mux.HandleFunc("POST /api/v1/tracking/session/{tripID}/stop", h.StopTracking)
}
