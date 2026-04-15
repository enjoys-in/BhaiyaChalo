package handler

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/dispatch", h.CreateDispatch)
	mux.HandleFunc("POST /api/v1/dispatch/respond", h.HandleDriverResponse)
	mux.HandleFunc("GET /api/v1/dispatch/offers/{offerID}", h.GetOfferStatus)
	mux.HandleFunc("POST /api/v1/dispatch/offers/{offerID}/expire", h.ExpireOffer)
}
