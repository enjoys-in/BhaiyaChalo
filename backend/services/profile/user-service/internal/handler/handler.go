package handler

import (
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/user-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/user-service/internal/ports"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/user-service/internal/repository"
)

type Handler struct {
	svc ports.UserService
}

func NewHandler(svc ports.UserService) *Handler {
	return &Handler{svc: svc}
}

// --- User CRUD ---

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.svc.CreateUser(r.Context(), req)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	successJSON(w, http.StatusCreated, "", resp)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := parseUUID(r, "id")
	if err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid user id")
		return
	}
	resp, err := h.svc.GetUser(r.Context(), id)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	successJSON(w, http.StatusOK, "", resp)
}

func (h *Handler) GetUserByPhone(w http.ResponseWriter, r *http.Request) {
	phone := r.URL.Query().Get("phone")
	if phone == "" {
		errorJSON(w, http.StatusBadRequest, "phone query param required")
		return
	}
	resp, err := h.svc.GetUserByPhone(r.Context(), phone)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	successJSON(w, http.StatusOK, "", resp)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := parseUUID(r, "id")
	if err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid user id")
		return
	}
	var req dto.UpdateUserRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.svc.UpdateUser(r.Context(), id, req)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	successJSON(w, http.StatusOK, "", resp)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := parseUUID(r, "id")
	if err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid user id")
		return
	}
	if err := h.svc.DeleteUser(r.Context(), id); err != nil {
		handleServiceError(w, err)
		return
	}
	successJSON(w, http.StatusNoContent, "", nil)
}

func (h *Handler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	id, err := parseUUID(r, "id")
	if err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid user id")
		return
	}
	resp, err := h.svc.GetUserProfile(r.Context(), id)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	successJSON(w, http.StatusOK, "", resp)
}

// --- Addresses ---

func (h *Handler) AddAddress(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUUID(r, "id")
	if err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid user id")
		return
	}
	var req dto.AddAddressRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.svc.AddAddress(r.Context(), userID, req)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	successJSON(w, http.StatusCreated, "", resp)
}

func (h *Handler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	addrID, err := parseUUID(r, "addressId")
	if err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid address id")
		return
	}
	var req dto.UpdateAddressRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.svc.UpdateAddress(r.Context(), addrID, req)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	successJSON(w, http.StatusOK, "", resp)
}

func (h *Handler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	addrID, err := parseUUID(r, "addressId")
	if err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid address id")
		return
	}
	if err := h.svc.DeleteAddress(r.Context(), addrID); err != nil {
		handleServiceError(w, err)
		return
	}
	successJSON(w, http.StatusNoContent, "", nil)
}

func (h *Handler) ListAddresses(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUUID(r, "id")
	if err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid user id")
		return
	}
	resp, err := h.svc.ListAddresses(r.Context(), userID)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	successJSON(w, http.StatusOK, "", resp)
}

// --- Emergency Contacts ---

func (h *Handler) AddEmergencyContact(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUUID(r, "id")
	if err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid user id")
		return
	}
	var req dto.AddEmergencyContactRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.svc.AddEmergencyContact(r.Context(), userID, req)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	successJSON(w, http.StatusCreated, "", resp)
}

func (h *Handler) ListEmergencyContacts(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUUID(r, "id")
	if err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid user id")
		return
	}
	resp, err := h.svc.ListEmergencyContacts(r.Context(), userID)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	successJSON(w, http.StatusOK, "", resp)
}

// --- helpers ---

func parseUUID(r *http.Request, key string) (uuid.UUID, error) {
	raw := r.PathValue(key)
	return uuid.Parse(raw)
}

func handleServiceError(w http.ResponseWriter, err error) {
	if errors.Is(err, repository.ErrNotFound) {
		errorJSON(w, http.StatusNotFound, "resource not found")
		return
	}
	errorJSON(w, http.StatusInternalServerError, "internal server error")
}
