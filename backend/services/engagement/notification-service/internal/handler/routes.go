package handler

import "net/http"

func (h *NotificationHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/notification", h.Create)
	mux.HandleFunc("GET /api/v1/notification/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/notification/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/notification/{id}", h.Delete)
}
