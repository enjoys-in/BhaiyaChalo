package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/iam-policy-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/iam-policy-service/internal/ports"
)

type Handler struct {
	svc ports.PolicyService
}

func NewHandler(svc ports.PolicyService) *Handler {
	return &Handler{svc: svc}
}

// --- Roles ---

func (h *Handler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateRoleRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.Name == "" {
		errorJSON(w, http.StatusBadRequest, "name is required")
		return
	}

	resp, err := h.svc.CreateRole(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "failed to create role")
		return
	}

	successJSON(w, http.StatusCreated, "", resp)
}

func (h *Handler) GetRole(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "role id is required")
		return
	}

	resp, err := h.svc.GetRole(r.Context(), id)
	if err != nil {
		errorJSON(w, http.StatusNotFound, "role not found")
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *Handler) ListRoles(w http.ResponseWriter, r *http.Request) {
	resp, err := h.svc.ListRoles(r.Context())
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "failed to list roles")
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

// --- Permissions ---

func (h *Handler) AssignPermission(w http.ResponseWriter, r *http.Request) {
	var req dto.AssignPermissionRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.RoleID == "" || req.PermissionID == "" {
		errorJSON(w, http.StatusBadRequest, "role_id and permission_id are required")
		return
	}

	if err := h.svc.AssignPermission(r.Context(), req); err != nil {
		errorJSON(w, http.StatusInternalServerError, "failed to assign permission")
		return
	}

	successJSON(w, http.StatusOK, "", map[string]string{"status": "assigned"})
}

func (h *Handler) GetPermissionsByRole(w http.ResponseWriter, r *http.Request) {
	roleID := r.PathValue("id")
	if roleID == "" {
		errorJSON(w, http.StatusBadRequest, "role id is required")
		return
	}

	resp, err := h.svc.GetPermissionsByRole(r.Context(), roleID)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "failed to get permissions")
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

// --- Policies ---

func (h *Handler) CreatePolicy(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePolicyRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.SubjectType == "" || req.SubjectID == "" || req.RoleID == "" {
		errorJSON(w, http.StatusBadRequest, "subject_type, subject_id, and role_id are required")
		return
	}

	resp, err := h.svc.CreatePolicy(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "failed to create policy")
		return
	}

	successJSON(w, http.StatusCreated, "", resp)
}

// --- Access Check ---

func (h *Handler) CheckAccess(w http.ResponseWriter, r *http.Request) {
	var req dto.CheckAccessRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.SubjectType == "" || req.SubjectID == "" || req.Resource == "" || req.Action == "" {
		errorJSON(w, http.StatusBadRequest, "subject_type, subject_id, resource, and action are required")
		return
	}

	resp, err := h.svc.CheckAccess(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "failed to check access")
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

// --- Health ---

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	successJSON(w, http.StatusOK, "", map[string]string{"status": "ok"})
}
