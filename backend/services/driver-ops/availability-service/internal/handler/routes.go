package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *AvailabilityHandler) {
	mux.HandleFunc("POST /api/v1/availability/online", h.GoOnline)
	mux.HandleFunc("POST /api/v1/availability/offline", h.GoOffline)
	mux.HandleFunc("PUT /api/v1/availability/trip-status", h.UpdateTripStatus)
	mux.HandleFunc("GET /api/v1/availability/count", h.CountOnline)
	mux.HandleFunc("GET /api/v1/availability/{driverID}", h.GetStatus)
}
