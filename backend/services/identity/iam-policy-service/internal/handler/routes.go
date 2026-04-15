package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/iam-policy-service/internal/ports"
)

func RegisterRoutes(mux *http.ServeMux, svc ports.PolicyService) {
	h := NewHandler(svc)

	// Roles
	mux.HandleFunc("POST /api/v1/roles", h.CreateRole)
	mux.HandleFunc("GET /api/v1/roles", h.ListRoles)
	mux.HandleFunc("GET /api/v1/roles/{id}", h.GetRole)

	// Permissions
	mux.HandleFunc("POST /api/v1/roles/permissions", h.AssignPermission)
	mux.HandleFunc("GET /api/v1/roles/{id}/permissions", h.GetPermissionsByRole)

	// Policies
	mux.HandleFunc("POST /api/v1/policies", h.CreatePolicy)

	// Access Check
	mux.HandleFunc("POST /api/v1/access/check", h.CheckAccess)

	// Health
	mux.HandleFunc("GET /healthz", h.HealthCheck)
}
