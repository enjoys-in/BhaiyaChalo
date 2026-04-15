package handler

import "net/http"

func (h *PayoutHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/payout", h.Create)
	mux.HandleFunc("GET /api/v1/payout/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/payout/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/payout/{id}", h.Delete)
}
