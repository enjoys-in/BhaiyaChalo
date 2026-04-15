package handler

import (
	"net/http"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/earnings-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/earnings-service/internal/ports"
)

type EarningsHandler struct {
	svc ports.EarningsService
}

func NewEarningsHandler(svc ports.EarningsService) *EarningsHandler {
	return &EarningsHandler{svc: svc}
}

// RecordEarning handles POST /api/v1/earnings
func (h *EarningsHandler) RecordEarning(w http.ResponseWriter, r *http.Request) {
	var req dto.RecordEarningRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.DriverID == "" || req.TripID == "" || req.FareAmount <= 0 {
		errorJSON(w, http.StatusBadRequest, "driver_id, trip_id and fare_amount are required")
		return
	}

	if req.Currency == "" {
		errorJSON(w, http.StatusBadRequest, "currency is required")
		return
	}

	resp, err := h.svc.RecordEarning(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusCreated, "", resp)
}

// GetEarnings handles GET /api/v1/earnings/{driverID}?from=...&to=...
func (h *EarningsHandler) GetEarnings(w http.ResponseWriter, r *http.Request) {
	driverID := pathParam(r, "driverID")
	if driverID == "" {
		errorJSON(w, http.StatusBadRequest, "driver_id is required")
		return
	}

	from, to := parseDateRange(r)

	resp, err := h.svc.GetEarnings(r.Context(), driverID, from, to)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

// GetDailySummary handles GET /api/v1/earnings/{driverID}/daily?date=...
func (h *EarningsHandler) GetDailySummary(w http.ResponseWriter, r *http.Request) {
	driverID := pathParam(r, "driverID")
	if driverID == "" {
		errorJSON(w, http.StatusBadRequest, "driver_id is required")
		return
	}

	date := parseDate(r.URL.Query().Get("date"))

	resp, err := h.svc.GetDailySummary(r.Context(), driverID, date)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

// GetWeeklySummary handles GET /api/v1/earnings/{driverID}/weekly?week_start=...
func (h *EarningsHandler) GetWeeklySummary(w http.ResponseWriter, r *http.Request) {
	driverID := pathParam(r, "driverID")
	if driverID == "" {
		errorJSON(w, http.StatusBadRequest, "driver_id is required")
		return
	}

	weekStart := parseDate(r.URL.Query().Get("week_start"))

	resp, err := h.svc.GetWeeklySummary(r.Context(), driverID, weekStart)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func parseDateRange(r *http.Request) (time.Time, time.Time) {
	from := parseDate(r.URL.Query().Get("from"))
	to := parseDate(r.URL.Query().Get("to"))
	if to.IsZero() {
		to = time.Now().UTC()
	}
	if from.IsZero() {
		from = to.AddDate(0, 0, -30)
	}
	return from, to
}

func parseDate(s string) time.Time {
	t, _ := time.Parse("2006-01-02", s)
	return t
}
