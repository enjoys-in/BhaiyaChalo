package handler

import "net/http"

func (h *FraudDetectionHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/fraud-detection", h.Create)
	mux.HandleFunc("GET /api/v1/fraud-detection/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/fraud-detection/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/fraud-detection/{id}", h.Delete)
}
