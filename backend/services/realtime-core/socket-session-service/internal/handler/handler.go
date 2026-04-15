package handler

import (
	"encoding/json"
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/ports"
)

type Handler struct {
	svc ports.SessionService
}

func NewHandler(svc ports.SessionService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.UserID == "" || req.Role == "" || req.DeviceID == "" || req.ServerID == "" {
		writeError(w, http.StatusBadRequest, "user_id, role, device_id and server_id are required")
		return
	}

	session, err := h.svc.Register(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, toSessionResponse(session))
}

func (h *Handler) Unregister(w http.ResponseWriter, r *http.Request) {
	var req dto.UnregisterSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.SessionID == "" || req.UserID == "" {
		writeError(w, http.StatusBadRequest, "session_id and user_id are required")
		return
	}

	if err := h.svc.Unregister(r.Context(), req); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "unregistered"})
}

func (h *Handler) GetByUser(w http.ResponseWriter, r *http.Request) {
	userID := extractPathParam(r, "userId")
	if userID == "" {
		writeError(w, http.StatusBadRequest, "user_id is required")
		return
	}

	sessions, err := h.svc.FindByUserID(r.Context(), userID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := toActiveSessionsResponse(sessions)
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) GetByServer(w http.ResponseWriter, r *http.Request) {
	serverID := extractPathParam(r, "serverId")
	if serverID == "" {
		writeError(w, http.StatusBadRequest, "server_id is required")
		return
	}

	sessions, err := h.svc.FindActiveByServer(r.Context(), serverID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := toActiveSessionsResponse(sessions)
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) CountActive(w http.ResponseWriter, r *http.Request) {
	count, err := h.svc.CountActive(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]int64{"active_count": count})
}
