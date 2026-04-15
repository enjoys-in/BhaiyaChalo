package handler

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/sessions/register", h.Register)
	mux.HandleFunc("POST /api/v1/sessions/unregister", h.Unregister)
	mux.HandleFunc("GET /api/v1/sessions/users/{userId}", h.GetByUser)
	mux.HandleFunc("GET /api/v1/sessions/servers/{serverId}", h.GetByServer)
	mux.HandleFunc("GET /api/v1/sessions/count", h.CountActive)
}
