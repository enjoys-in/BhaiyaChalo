package handler

import (
	"encoding/json"
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/promo-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/promo-service/internal/ports"
)

type PromoHandler struct {
	svc ports.PromoService
}

func NewPromoHandler(svc ports.PromoService) *PromoHandler {
	return &PromoHandler{svc: svc}
}

func (h *PromoHandler) CreatePromo(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePromoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := validateCreatePromoRequest(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.svc.Create(r.Context(), &req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, resp)
}

func (h *PromoHandler) ApplyPromo(w http.ResponseWriter, r *http.Request) {
	var req dto.ApplyPromoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Code == "" || req.UserID == "" || req.CityID == "" {
		writeError(w, http.StatusBadRequest, "code, user_id, and city_id are required")
		return
	}
	if req.BookingAmount <= 0 {
		writeError(w, http.StatusBadRequest, "booking_amount must be greater than 0")
		return
	}

	resp, err := h.svc.Apply(r.Context(), &req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *PromoHandler) ValidatePromo(w http.ResponseWriter, r *http.Request) {
	var req dto.ValidatePromoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Code == "" || req.UserID == "" || req.CityID == "" {
		writeError(w, http.StatusBadRequest, "code, user_id, and city_id are required")
		return
	}

	resp, err := h.svc.Validate(r.Context(), &req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *PromoHandler) ListActivePromos(w http.ResponseWriter, r *http.Request) {
	cityID := r.URL.Query().Get("city_id")

	resp, err := h.svc.ListActive(r.Context(), cityID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *PromoHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
