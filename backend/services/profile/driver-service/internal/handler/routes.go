package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *DriverHandler) {
	mux.HandleFunc("POST /api/v1/drivers", h.Create)
	mux.HandleFunc("GET /api/v1/drivers/{id}", h.GetByID)
	mux.HandleFunc("GET /api/v1/drivers/phone/{phone}", h.GetByPhone)
	mux.HandleFunc("PUT /api/v1/drivers/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/drivers/{id}", h.Delete)
	mux.HandleFunc("GET /api/v1/drivers/city/{cityId}", h.ListByCityID)

	mux.HandleFunc("GET /api/v1/drivers/{id}/preferences", h.GetPreference)
	mux.HandleFunc("PUT /api/v1/drivers/{id}/preferences", h.UpdatePreference)
}
