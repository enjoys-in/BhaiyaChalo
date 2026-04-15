package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/search-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/search-service/internal/ports"
)

type SearchHandler struct {
	svc ports.SearchService
}

func NewSearchHandler(svc ports.SearchService) *SearchHandler {
	return &SearchHandler{svc: svc}
}

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	var req dto.SearchRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.CityID == "" || req.UserID == "" {
		errorJSON(w, http.StatusBadRequest, "city_id and user_id are required")
		return
	}

	resp, err := h.svc.Search(r.Context(), &req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", resp)
}
