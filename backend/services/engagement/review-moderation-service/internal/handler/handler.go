package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/review-moderation-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/engagement/review-moderation-service/internal/ports"
)

type ReviewModerationHandler struct {
	svc ports.ReviewModerationService
}

func NewReviewModerationHandler(svc ports.ReviewModerationService) *ReviewModerationHandler {
	return &ReviewModerationHandler{svc: svc}
}

func (h *ReviewModerationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateReviewRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Create(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusCreated, "reviewModeration created successfully", resp)
}

func (h *ReviewModerationHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	successJSON(w, http.StatusOK, "reviewModeration fetched successfully", resp)
}

func (h *ReviewModerationHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateReviewRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.svc.Update(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "reviewModeration updated successfully", resp)
}

func (h *ReviewModerationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "id is required")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJSON(w, http.StatusOK, "reviewModeration deleted successfully", nil)
}
