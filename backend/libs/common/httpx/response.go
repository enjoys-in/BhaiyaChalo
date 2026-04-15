package httpx

import (
	"encoding/json"
	"net/http"

	apperr "github.com/enjoys-in/BhaiyaChalo/libs/common/errors"
)

// Response is the standard JSON envelope for all APIs (user, admin, driver).
// Format: {"success": bool, "message": "...", "result": ... }
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// JSON writes a success response.
func JSON(w http.ResponseWriter, status int, message string, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{Success: true, Message: message, Result: result})
}

// Error writes an error response.
func Error(w http.ResponseWriter, err error) {
	if ae, ok := err.(*apperr.AppError); ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(ae.Code)
		json.NewEncoder(w).Encode(Response{Success: false, Message: ae.Message, Result: nil})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Response{Success: false, Message: "internal server error", Result: nil})
}

// ErrorWithStatus writes an error response with a specific HTTP status.
func ErrorWithStatus(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{Success: false, Message: message, Result: nil})
}

// Decode reads JSON body into v and returns an AppError on failure.
func Decode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return apperr.NewBadRequest("invalid request body")
	}
	return nil
}
