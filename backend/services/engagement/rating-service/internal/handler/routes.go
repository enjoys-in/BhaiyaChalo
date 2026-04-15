package handler

import "net/http"

func (h *RatingHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/rating", h.Create)
	mux.HandleFunc("GET /api/v1/rating/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/rating/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/rating/{id}", h.Delete)
}
