package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *DocumentHandler) {
	mux.HandleFunc("POST /api/v1/documents", h.Upload)
	mux.HandleFunc("GET /api/v1/documents/pending", h.ListPending)
	mux.HandleFunc("GET /api/v1/documents/expiring", h.ListExpiring)
	mux.HandleFunc("GET /api/v1/documents/{id}", h.GetByID)
	mux.HandleFunc("POST /api/v1/documents/{id}/review", h.Review)
	mux.HandleFunc("GET /api/v1/{ownerType}/{ownerId}/documents", h.GetByOwner)
	mux.HandleFunc("GET /healthz", h.HealthCheck)
}
