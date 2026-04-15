package handler

import "net/http"

func (h *SupportTicketHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/support-ticket", h.Create)
	mux.HandleFunc("GET /api/v1/support-ticket/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/support-ticket/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/support-ticket/{id}", h.Delete)
}
