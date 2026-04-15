package handler

import "net/http"

func (h *AuditLogHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/audit-log", h.Create)
	mux.HandleFunc("GET /api/v1/audit-log/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/audit-log/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/audit-log/{id}", h.Delete)
}
