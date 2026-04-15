package handler

import "net/http"

func (h *EscalationHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/escalation", h.Create)
	mux.HandleFunc("GET /api/v1/escalation/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/escalation/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/escalation/{id}", h.Delete)
}
