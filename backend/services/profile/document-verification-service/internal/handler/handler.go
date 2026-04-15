package handler

import (
	"net/http"
	"strconv"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/document-verification-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/document-verification-service/internal/ports"
)

type DocumentHandler struct {
	svc ports.DocumentService
}

func NewDocumentHandler(svc ports.DocumentService) *DocumentHandler {
	return &DocumentHandler{svc: svc}
}

func (h *DocumentHandler) Upload(w http.ResponseWriter, r *http.Request) {
	var req dto.UploadDocumentRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.OwnerID == "" || req.OwnerType == "" || req.DocType == "" || req.FileURL == "" {
		errorJSON(w, http.StatusBadRequest, "owner_id, owner_type, doc_type and file_url are required")
		return
	}

	resp, err := h.svc.Upload(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusCreated, "", resp)
}

func (h *DocumentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "document id is required")
		return
	}

	resp, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		errorJSON(w, http.StatusNotFound, "document not found")
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *DocumentHandler) GetByOwner(w http.ResponseWriter, r *http.Request) {
	ownerID := r.PathValue("ownerId")
	ownerType := r.PathValue("ownerType")

	if ownerID == "" || ownerType == "" {
		errorJSON(w, http.StatusBadRequest, "owner_id and owner_type are required")
		return
	}

	resp, err := h.svc.GetByOwner(r.Context(), ownerID, ownerType)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *DocumentHandler) Review(w http.ResponseWriter, r *http.Request) {
	var req dto.ReviewDocumentRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "document id is required")
		return
	}
	req.DocumentID = id

	if req.ReviewerID == "" || req.Status == "" {
		errorJSON(w, http.StatusBadRequest, "reviewer_id and status are required")
		return
	}

	resp, err := h.svc.Review(r.Context(), req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *DocumentHandler) ListPending(w http.ResponseWriter, r *http.Request) {
	page := queryInt(r, "page", 1)
	perPage := queryInt(r, "per_page", 20)

	resp, err := h.svc.ListPending(r.Context(), page, perPage)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *DocumentHandler) ListExpiring(w http.ResponseWriter, r *http.Request) {
	daysAhead := queryInt(r, "days", 30)

	resp, err := h.svc.ListExpiring(r.Context(), daysAhead)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *DocumentHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	successJSON(w, http.StatusOK, "", map[string]string{"status": "ok"})
}

func queryInt(r *http.Request, key string, fallback int) int {
	v := r.URL.Query().Get(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}
