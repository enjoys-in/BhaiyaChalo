package handler

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/connections", h.RegisterConnection)
	mux.HandleFunc("DELETE /api/v1/connections/{connectionId}", h.RemoveConnection)
	mux.HandleFunc("GET /api/v1/connections/users/{userId}/locate", h.LocateUser)
	mux.HandleFunc("GET /api/v1/connections/nodes/{nodeId}/health", h.GetNodeHealth)
}
