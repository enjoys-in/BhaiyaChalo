package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/delivery-ack-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/delivery-ack-service/internal/ports"
)

type DeliveryAckHandler struct {
	svc ports.DeliveryAckService
}

func NewDeliveryAckHandler(svc ports.DeliveryAckService) *DeliveryAckHandler {
	return &DeliveryAckHandler{svc: svc}
}

func (h *DeliveryAckHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateDeliveryAckRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusCreated, "deliveryAck created successfully", resp)
}

func (h *DeliveryAckHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	resp, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "deliveryAck fetched successfully", resp)
}

func (h *DeliveryAckHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateDeliveryAckRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "deliveryAck updated successfully", resp)
}

func (h *DeliveryAckHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "deliveryAck deleted successfully", nil)
}
