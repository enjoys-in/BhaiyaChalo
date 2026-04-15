package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	// Users
	mux.HandleFunc("POST /api/v1/users", h.CreateUser)
	mux.HandleFunc("GET /api/v1/users/search", h.GetUserByPhone)
	mux.HandleFunc("GET /api/v1/users/{id}", h.GetUser)
	mux.HandleFunc("PATCH /api/v1/users/{id}", h.UpdateUser)
	mux.HandleFunc("DELETE /api/v1/users/{id}", h.DeleteUser)
	mux.HandleFunc("GET /api/v1/users/{id}/profile", h.GetUserProfile)

	// Addresses
	mux.HandleFunc("POST /api/v1/users/{id}/addresses", h.AddAddress)
	mux.HandleFunc("GET /api/v1/users/{id}/addresses", h.ListAddresses)
	mux.HandleFunc("PATCH /api/v1/users/{id}/addresses/{addressId}", h.UpdateAddress)
	mux.HandleFunc("DELETE /api/v1/users/{id}/addresses/{addressId}", h.DeleteAddress)

	// Emergency Contacts
	mux.HandleFunc("POST /api/v1/users/{id}/emergency-contacts", h.AddEmergencyContact)
	mux.HandleFunc("GET /api/v1/users/{id}/emergency-contacts", h.ListEmergencyContacts)
}
