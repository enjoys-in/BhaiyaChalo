package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *PricingHandler) {
	mux.HandleFunc("POST /api/v1/pricing/estimate", h.Estimate)
	mux.HandleFunc("GET /api/v1/pricing/rules", h.GetRules)
	mux.HandleFunc("POST /api/v1/pricing/rules", h.CreateRule)
	mux.HandleFunc("PUT /api/v1/pricing/rules/{id}", h.UpdateRule)
}
