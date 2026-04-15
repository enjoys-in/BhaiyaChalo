package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *BookingHandler) {
	mux.HandleFunc("POST /api/v1/bookings", h.CreateBooking)
	mux.HandleFunc("GET /api/v1/bookings/{id}", h.GetBooking)
	mux.HandleFunc("POST /api/v1/bookings/cancel", h.CancelBooking)
	mux.HandleFunc("PUT /api/v1/bookings/status", h.UpdateStatus)
	mux.HandleFunc("POST /api/v1/bookings/assign-driver", h.AssignDriver)
}
