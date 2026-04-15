package handler

import (
	"encoding/json"
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/connection-manager-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/connection-manager-service/internal/ports"
)

type Handler struct {
	svc ports.ConnectionManager
}

func NewHandler(svc ports.ConnectionManager) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterConnection(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterConnectionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.UserID == "" || req.ServerNode == "" || req.Protocol == "" {
		writeError(w, http.StatusBadRequest, "user_id, server_node and protocol are required")
		return
	}

	conn, err := h.svc.Register(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, toConnectionResponse(conn))
}

func (h *Handler) RemoveConnection(w http.ResponseWriter, r *http.Request) {
	connID := extractPathParam(r, "connectionId")
	if connID == "" {
		writeError(w, http.StatusBadRequest, "connection_id is required")
		return
	}

	if err := h.svc.Remove(r.Context(), connID); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "removed"})
}

func (h *Handler) LocateUser(w http.ResponseWriter, r *http.Request) {
	userID := extractPathParam(r, "userId")
	if userID == "" {
		writeError(w, http.StatusBadRequest, "user_id is required")
		return
	}

	conn, err := h.svc.LocateUser(r.Context(), userID)
	if err != nil {
		writeError(w, http.StatusNotFound, "user not connected")
		return
	}

	writeJSON(w, http.StatusOK, dto.UserLocationResponse{
		UserID:     conn.UserID,
		ServerNode: conn.ServerNode,
		Connected:  true,
	})
}

func (h *Handler) GetNodeHealth(w http.ResponseWriter, r *http.Request) {
	nodeID := extractPathParam(r, "nodeId")
	if nodeID == "" {
		writeError(w, http.StatusBadRequest, "node_id is required")
		return
	}

	status, err := h.svc.GetNodeHealth(r.Context(), nodeID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, toNodeStatusResponse(status))
}
