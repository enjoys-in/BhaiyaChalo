package handler

import "net/http"

func (h *ReconciliationHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/reconciliation", h.Create)
	mux.HandleFunc("GET /api/v1/reconciliation/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/reconciliation/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/reconciliation/{id}", h.Delete)
}
