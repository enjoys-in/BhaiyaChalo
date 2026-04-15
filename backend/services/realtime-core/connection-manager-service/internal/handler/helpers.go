package handler

import (
	"encoding/json"
	"net/http"
)

type apiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func successJSON(w http.ResponseWriter, status int, message string, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(apiResponse{Success: true, Message: message, Result: result})
}

func errorJSON(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(apiResponse{Success: false, Message: msg, Result: nil})
}

func decodeJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	errorJSON(w, status, msg)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	successJSON(w, status, "", data)
}

func extractPathParam(r *http.Request, key string) string {
	return r.PathValue(key)
}

func pathParam(r *http.Request, key string) string {
	return r.PathValue(key)
}

func toConnectionResponse(data interface{}) interface{} {
	// TODO: implement response mapping
	return data
}

func toNodeStatusResponse(data interface{}) interface{} {
	// TODO: implement response mapping
	return data
}
