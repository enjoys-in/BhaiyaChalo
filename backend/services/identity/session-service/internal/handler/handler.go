package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/session-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/session-service/internal/ports"
)

type SessionHandler struct {
	svc ports.SessionService
}

func NewSessionHandler(svc ports.SessionService) *SessionHandler {
	return &SessionHandler{svc: svc}
}

func (h *SessionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateSessionRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.UserID == "" || req.Role == "" {
		errorJSON(w, http.StatusBadRequest, "user_id and role are required")
		return
	}

	resp, err := h.svc.Create(r.Context(), &req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "failed to create session")
		return
	}

	successJSON(w, http.StatusCreated, "", resp)
}

func (h *SessionHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "session id is required")
		return
	}

	resp, err := h.svc.GetByID(r.Context(), &dto.GetSessionRequest{SessionID: id})
	if err != nil {
		errorJSON(w, http.StatusNotFound, "session not found")
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *SessionHandler) Invalidate(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "session id is required")
		return
	}

	err := h.svc.Invalidate(r.Context(), &dto.InvalidateRequest{SessionID: id})
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "failed to invalidate session")
		return
	}

	successJSON(w, http.StatusOK, "", map[string]string{"status": "invalidated"})
}

func (h *SessionHandler) InvalidateAll(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	if userID == "" {
		errorJSON(w, http.StatusBadRequest, "user id is required")
		return
	}

	err := h.svc.InvalidateAll(r.Context(), &dto.InvalidateAllRequest{UserID: userID})
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "failed to invalidate sessions")
		return
	}

	successJSON(w, http.StatusOK, "", map[string]string{"status": "all sessions invalidated"})
}
