package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/template-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/engagement/template-service/internal/ports"
)

type TemplateHandler struct {
	svc ports.TemplateService
}

func NewTemplateHandler(svc ports.TemplateService) *TemplateHandler {
	return &TemplateHandler{svc: svc}
}

func (h *TemplateHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTemplateRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusCreated, "template created successfully", resp)
}

func (h *TemplateHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	successJSON(w, http.StatusOK, "template fetched successfully", resp)
}

func (h *TemplateHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateTemplateRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "template updated successfully", resp)
}

func (h *TemplateHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "template deleted successfully", nil)
}
