package handler

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/routes", h.PlanRoute)
	mux.HandleFunc("GET /api/v1/routes/{bookingId}", h.GetRoute)
}
