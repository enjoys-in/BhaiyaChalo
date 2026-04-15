package handler

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/trips/{tripId}/stops", h.AddStop)
	mux.HandleFunc("DELETE /api/v1/trips/{tripId}/stops/{stopId}", h.RemoveStop)
	mux.HandleFunc("PUT /api/v1/trips/{tripId}/stops/reorder", h.ReorderStops)
	mux.HandleFunc("PATCH /api/v1/trips/{tripId}/stops/{stopId}/status", h.UpdateStopStatus)
	mux.HandleFunc("GET /api/v1/trips/{tripId}/stops", h.GetStopsByTrip)
}
