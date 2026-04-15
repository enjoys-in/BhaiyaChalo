package handler

import "net/http"

func (h *RiskScoringHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/risk-scoring", h.Create)
	mux.HandleFunc("GET /api/v1/risk-scoring/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/risk-scoring/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/risk-scoring/{id}", h.Delete)
}
