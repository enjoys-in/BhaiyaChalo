package handler

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/fares/calculate", h.CalculateFare)
	mux.HandleFunc("POST /api/v1/fares/recalculate", h.RecalculateFare)
	mux.HandleFunc("GET /api/v1/fares/{bookingId}", h.GetBreakdown)
}
