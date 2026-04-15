package handler

import "net/http"

func (h *TemplateHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/template", h.Create)
	mux.HandleFunc("GET /api/v1/template/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/template/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/template/{id}", h.Delete)
}
