package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *SearchHandler) {
	mux.HandleFunc("POST /api/v1/search", h.Search)
}
