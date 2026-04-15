package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/support/support-ticket-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/support/support-ticket-service/internal/ports"
)

type SupportTicketHandler struct {
	svc ports.SupportTicketService
}

func NewSupportTicketHandler(svc ports.SupportTicketService) *SupportTicketHandler {
	return &SupportTicketHandler{svc: svc}
}

func (h *SupportTicketHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTicketRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusCreated, "supportTicket created successfully", resp)
}

func (h *SupportTicketHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	successJSON(w, http.StatusOK, "supportTicket fetched successfully", resp)
}

func (h *SupportTicketHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateTicketRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "supportTicket updated successfully", resp)
}

func (h *SupportTicketHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "supportTicket deleted successfully", nil)
}
