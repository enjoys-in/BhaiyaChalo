package handler

import "net/http"

func (h *PromoHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.HealthCheck)
	mux.HandleFunc("POST /api/v1/promos", h.CreatePromo)
	mux.HandleFunc("POST /api/v1/promos/apply", h.ApplyPromo)
	mux.HandleFunc("POST /api/v1/promos/validate", h.ValidatePromo)
	mux.HandleFunc("GET /api/v1/promos", h.ListActivePromos)
}
