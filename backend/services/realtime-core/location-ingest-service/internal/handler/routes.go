package handler

import "net/http"

func (h *LocationIngestHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/location-ingest", h.Create)
	mux.HandleFunc("GET /api/v1/location-ingest/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/location-ingest/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/location-ingest/{id}", h.Delete)
}
