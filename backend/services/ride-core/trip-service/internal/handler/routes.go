package handler

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/trips", h.CreateTrip)
	mux.HandleFunc("GET /api/v1/trips/{tripId}", h.GetTrip)
	mux.HandleFunc("PATCH /api/v1/trips/{tripId}/status", h.UpdateTripStatus)
	mux.HandleFunc("GET /api/v1/trips/{tripId}/timeline", h.GetTimeline)
	mux.HandleFunc("GET /api/v1/users/{userId}/trips", h.ListByUser)
	mux.HandleFunc("GET /api/v1/drivers/{driverId}/trips", h.ListByDriver)
}
