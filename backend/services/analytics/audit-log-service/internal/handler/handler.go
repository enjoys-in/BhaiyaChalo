package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/audit-log-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/analytics/audit-log-service/internal/ports"
)

type AuditLogHandler struct {
	svc ports.AuditLogService
}

func NewAuditLogHandler(svc ports.AuditLogService) *AuditLogHandler {
	return &AuditLogHandler{svc: svc}
}

func (h *AuditLogHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateAuditEntryRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusCreated, "auditLog created successfully", resp)
}

func (h *AuditLogHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	successJSON(w, http.StatusOK, "auditLog fetched successfully", resp)
}

func (h *AuditLogHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateAuditEntryRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "auditLog updated successfully", resp)
}

func (h *AuditLogHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "auditLog deleted successfully", nil)
}
