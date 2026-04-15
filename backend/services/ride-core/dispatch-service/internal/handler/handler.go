package handler

import (
	"encoding/json"
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/ports"
)

type Handler struct {
	svc ports.DispatchService
}

func NewHandler(svc ports.DispatchService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateDispatch(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateDispatchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := validateCreateDispatchRequest(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	offer, err := h.svc.Dispatch(r.Context(), &req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, toDispatchResponse(offer))
}

func (h *Handler) HandleDriverResponse(w http.ResponseWriter, r *http.Request) {
	var req dto.DriverResponseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.OfferID == "" {
		writeError(w, http.StatusBadRequest, "offer_id is required")
		return
	}

	offer, err := h.svc.HandleDriverResponse(r.Context(), &req)
	if err != nil {
		writeError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, toOfferStatusResponse(offer))
}

func (h *Handler) GetOfferStatus(w http.ResponseWriter, r *http.Request) {
	offerID := extractPathParam(r, "offerID")
	if offerID == "" {
		writeError(w, http.StatusBadRequest, "offer_id is required")
		return
	}

	offer, err := h.svc.HandleDriverResponse(r.Context(), &dto.DriverResponseRequest{OfferID: offerID})
	if err != nil {
		// Fallback: just try to read the offer status directly
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, toOfferStatusResponse(offer))
}

func (h *Handler) ExpireOffer(w http.ResponseWriter, r *http.Request) {
	offerID := extractPathParam(r, "offerID")
	if offerID == "" {
		writeError(w, http.StatusBadRequest, "offer_id is required")
		return
	}

	if err := h.svc.ExpireOffer(r.Context(), offerID); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "expired"})
}
