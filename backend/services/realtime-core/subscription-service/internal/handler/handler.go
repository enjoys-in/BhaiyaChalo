package handler

import (
	"encoding/json"
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/subscription-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/subscription-service/internal/ports"
)

type Handler struct {
	svc ports.SubscriptionService
}

func NewHandler(svc ports.SubscriptionService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	var req dto.SubscribeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.UserID == "" || req.Channel == "" || req.Topic == "" {
		writeError(w, http.StatusBadRequest, "user_id, channel and topic are required")
		return
	}

	sub, err := h.svc.Subscribe(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, toSubscriptionResponse(sub))
}

func (h *Handler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	var req dto.UnsubscribeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.UserID == "" || req.SubscriptionID == "" {
		writeError(w, http.StatusBadRequest, "user_id and subscription_id are required")
		return
	}

	if err := h.svc.Unsubscribe(r.Context(), req); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "unsubscribed"})
}

func (h *Handler) GetByUser(w http.ResponseWriter, r *http.Request) {
	userID := extractPathParam(r, "userId")
	if userID == "" {
		writeError(w, http.StatusBadRequest, "user_id is required")
		return
	}

	subs, err := h.svc.FindByUser(r.Context(), userID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, toSubscriptionListResponse(subs))
}

func (h *Handler) GetByChannel(w http.ResponseWriter, r *http.Request) {
	channel := extractPathParam(r, "channel")
	topic := r.URL.Query().Get("topic")

	if channel == "" {
		writeError(w, http.StatusBadRequest, "channel is required")
		return
	}

	subs, err := h.svc.FindByChannel(r.Context(), channel, topic)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, toSubscriptionListResponse(subs))
}

func (h *Handler) CountByChannel(w http.ResponseWriter, r *http.Request) {
	channel := extractPathParam(r, "channel")
	if channel == "" {
		writeError(w, http.StatusBadRequest, "channel is required")
		return
	}

	count, err := h.svc.CountByChannel(r.Context(), channel)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]int64{"count": count})
}
