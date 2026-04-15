package handler

import "net/http"

func (h *IncentiveHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/incentive", h.Create)
	mux.HandleFunc("GET /api/v1/incentive/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/incentive/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/incentive/{id}", h.Delete)
}
