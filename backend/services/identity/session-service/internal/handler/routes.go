package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *SessionHandler) {
	mux.HandleFunc("POST /api/v1/sessions", h.Create)
	mux.HandleFunc("GET /api/v1/sessions/{id}", h.Get)
	mux.HandleFunc("DELETE /api/v1/sessions/{id}", h.Invalidate)
	mux.HandleFunc("DELETE /api/v1/sessions/user/{userId}", h.InvalidateAll)
}
