package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *EarningsHandler) {
	mux.HandleFunc("POST /api/v1/earnings", h.RecordEarning)
	mux.HandleFunc("GET /api/v1/earnings/{driverID}", h.GetEarnings)
	mux.HandleFunc("GET /api/v1/earnings/{driverID}/daily", h.GetDailySummary)
	mux.HandleFunc("GET /api/v1/earnings/{driverID}/weekly", h.GetWeeklySummary)
}
