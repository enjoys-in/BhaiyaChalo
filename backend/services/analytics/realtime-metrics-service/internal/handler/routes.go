package handler

import "net/http"

func (h *RealtimeMetricsHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/realtime-metrics", h.Create)
	mux.HandleFunc("GET /api/v1/realtime-metrics/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/realtime-metrics/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/realtime-metrics/{id}", h.Delete)
}
