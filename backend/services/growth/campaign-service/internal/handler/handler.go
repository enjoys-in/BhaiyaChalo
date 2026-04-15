package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/campaign-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/campaign-service/internal/ports"
)

type CampaignHandler struct {
	svc ports.CampaignService
}

func NewCampaignHandler(svc ports.CampaignService) *CampaignHandler {
	return &CampaignHandler{svc: svc}
}

func (h *CampaignHandler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCampaignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := validateCreateCampaignRequest(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.svc.Create(r.Context(), &req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, resp)
}

func (h *CampaignHandler) LaunchCampaign(w http.ResponseWriter, r *http.Request) {
	campaignID := extractPathParam(r.URL.Path, "/api/v1/campaigns/", "/launch")
	if campaignID == "" {
		writeError(w, http.StatusBadRequest, "campaign_id is required")
		return
	}

	resp, err := h.svc.Launch(r.Context(), campaignID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *CampaignHandler) PauseCampaign(w http.ResponseWriter, r *http.Request) {
	campaignID := extractPathParam(r.URL.Path, "/api/v1/campaigns/", "/pause")
	if campaignID == "" {
		writeError(w, http.StatusBadRequest, "campaign_id is required")
		return
	}

	resp, err := h.svc.Pause(r.Context(), campaignID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *CampaignHandler) GetCampaignStats(w http.ResponseWriter, r *http.Request) {
	campaignID := extractPathParam(r.URL.Path, "/api/v1/campaigns/", "/stats")
	if campaignID == "" {
		writeError(w, http.StatusBadRequest, "campaign_id is required")
		return
	}

	resp, err := h.svc.GetStats(r.Context(), campaignID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *CampaignHandler) ListCampaigns(w http.ResponseWriter, r *http.Request) {
	cityID := r.URL.Query().Get("city_id")

	resp, err := h.svc.List(r.Context(), cityID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *CampaignHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func extractPathParam(path, prefix, suffix string) string {
	path = strings.TrimPrefix(path, prefix)
	path = strings.TrimSuffix(path, suffix)
	return path
}
