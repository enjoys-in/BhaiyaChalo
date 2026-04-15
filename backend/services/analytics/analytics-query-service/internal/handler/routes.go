package handler

import "net/http"

func (h *AnalyticsQueryHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/analytics-query", h.Create)
	mux.HandleFunc("GET /api/v1/analytics-query/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/analytics-query/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/analytics-query/{id}", h.Delete)
}
