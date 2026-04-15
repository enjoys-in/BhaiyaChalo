package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *VehicleHandler) {
	mux.HandleFunc("POST /api/v1/vehicles", h.Create)
	mux.HandleFunc("GET /api/v1/vehicles", h.ListByType)
	mux.HandleFunc("GET /api/v1/vehicles/{id}", h.GetByID)
	mux.HandleFunc("PUT /api/v1/vehicles/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/vehicles/{id}", h.Delete)
	mux.HandleFunc("GET /api/v1/drivers/{driverId}/vehicles", h.GetByDriverID)
	mux.HandleFunc("GET /healthz", h.HealthCheck)
}
