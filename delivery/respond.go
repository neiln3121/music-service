package delivery

import (
	"encoding/json"
	"net/http"
)

type appError struct {
	Message string
	Cause   error
}

func newAppError(message string, cause error) *appError {
	return &appError{
		Message: message,
		Cause:   cause,
	}
}

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
