package handler

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/subscriptions", h.Subscribe)
	mux.HandleFunc("POST /api/v1/subscriptions/unsubscribe", h.Unsubscribe)
	mux.HandleFunc("GET /api/v1/subscriptions/users/{userId}", h.GetByUser)
	mux.HandleFunc("GET /api/v1/subscriptions/channels/{channel}", h.GetByChannel)
	mux.HandleFunc("GET /api/v1/subscriptions/channels/{channel}/count", h.CountByChannel)
}
