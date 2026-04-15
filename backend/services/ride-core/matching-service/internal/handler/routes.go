package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *MatchingHandler) {
	mux.HandleFunc("POST /api/v1/matching/find", h.FindDrivers)
	mux.HandleFunc("POST /api/v1/matching/assign", h.AssignDriver)
}
