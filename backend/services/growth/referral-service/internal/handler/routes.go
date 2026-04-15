package handler

import "net/http"

func (h *ReferralHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.HealthCheck)
	mux.HandleFunc("POST /api/v1/referrals/code", h.GenerateCode)
	mux.HandleFunc("POST /api/v1/referrals/apply", h.ApplyReferral)
	mux.HandleFunc("POST /api/v1/referrals/{id}/complete", h.CompleteReferral)
	mux.HandleFunc("POST /api/v1/referrals/reward", h.ClaimReward)
	mux.HandleFunc("GET /api/v1/referrals/stats", h.GetStats)
}
