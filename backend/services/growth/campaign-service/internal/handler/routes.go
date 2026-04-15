package handler

import "net/http"

func (h *CampaignHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.HealthCheck)
	mux.HandleFunc("POST /api/v1/campaigns", h.CreateCampaign)
	mux.HandleFunc("GET /api/v1/campaigns", h.ListCampaigns)
	mux.HandleFunc("POST /api/v1/campaigns/{id}/launch", h.LaunchCampaign)
	mux.HandleFunc("POST /api/v1/campaigns/{id}/pause", h.PauseCampaign)
	mux.HandleFunc("GET /api/v1/campaigns/{id}/stats", h.GetCampaignStats)
}
