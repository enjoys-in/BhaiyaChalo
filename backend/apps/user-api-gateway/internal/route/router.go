package route

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/apps/user-api-gateway/internal/handler"
	"github.com/enjoys-in/BhaiyaChalo/apps/user-api-gateway/internal/middleware"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("GET /healthz", handler.HealthCheck)
	mux.HandleFunc("GET /readyz", handler.ReadyCheck)

	// TODO: register proxy routes to downstream services

	// Apply middleware chain
	var h http.Handler = mux
	h = middleware.CORSMiddleware(h)
	h = middleware.RateLimitMiddleware(h)
	h = middleware.AuthMiddleware(h)

	return h
}
